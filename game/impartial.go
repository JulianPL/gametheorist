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
