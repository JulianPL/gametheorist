package game

import (
	"fmt"
	"gametheorist/collection"
	"gametheorist/number"
)

type Impartial struct {
	value number.Nimber
}

func MakeGameImpartial(s string) Impartial {
	value := number.MakeNimber(s)
	return Impartial{value}
}

func (x Impartial) String() string {
	return fmt.Sprintf(x.value.String())
}

func (x Impartial) Add(y Impartial) Impartial {
	return Impartial{x.value.Add(y.value)}
}

func (x Impartial) Increment() Impartial {
	value := x.value.Increment()
	return Impartial{value}
}

func MinimalExcluded(children collection.Set[Impartial]) Impartial {
	mex := MakeGameImpartial("0")
	for children.Contains(mex) {
		mex = mex.Increment()
	}
	return mex
}

// NimSequenceFromSubtractionSet calculates the first length nim-values for a nim-style game,
// in which the number of taken elements has to be in the subtractionSet.
// It uses a basic Grundy-Scale and the minimal excluded property.
// Further improvements like Ferguson's Pairing Property (p.86 winning ways) are not implemented.
func NimSequenceFromSubtractionSet(subtractionSet collection.Set[int], length int) []Impartial {
	games := make([]Impartial, length)
	for i := 0; i < length; i++ {
		children := collection.MakeSet[Impartial]()
		for move := range subtractionSet {
			if i-move < 0 {
				continue
			}
			children.Add(games[i-move])
		}
		games[i] = MinimalExcluded(children)
	}
	return games
}

// NimSequenceFromTakeAndBreakCode calculates the first length nim-values for a nim-style game,
// in which the number of taken elements and the possible counts of subheaps is given by the takeAndBreakCode.
// It uses the minimal excluded property on all suitable subpartitions.
// Todo Implement improvements after p.91 of winning ways (if applicable) or at least refactor
func NimSequenceFromTakeAndBreakCode(takeAndBreakCode []number.BigInt, length int) []Impartial {
	games := make([]Impartial, length)
	for elementCount := 0; elementCount < length; elementCount++ {
		children := collection.MakeSet[Impartial]()
		for index, takeAndBreakValue := range takeAndBreakCode {
			removed := index + 1
			heapCount := 0
			for takeAndBreakValue.Greater(number.MakeBigInt("0")) {
				if takeAndBreakValue.IsOdd() {
					partitions := collection.MakeOrderedPartitions(elementCount-removed, heapCount)
					for _, partition := range partitions {
						nimber := MakeGameImpartial("0")
						for _, heap := range partition {
							nimber = nimber.Add(games[heap])
						}
						children.Add(nimber)
					}
				}
				heapCount++
				takeAndBreakValue = takeAndBreakValue.Div(number.MakeBigInt("2"))
			}
		}
		games[elementCount] = MinimalExcluded(children)
	}
	return games
}
