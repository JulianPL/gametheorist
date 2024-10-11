package game

import (
	"fmt"
	"gametheorist/collection"
	"gametheorist/number"
)

type Rat struct {
	value number.BigRat
}

func MakeGameRat(s string) Rat {
	value := number.MakeBigRat(s)
	return Rat{value}
}

func (x Rat) String() string {
	return fmt.Sprintf(x.value.String())
}

func (x Rat) Less(y Rat) bool {
	return x.value.Less(y.value)
}

func (x Rat) Greater(y Rat) bool {
	return x.value.Greater(y.value)
}

func (x Rat) IsInt() bool {
	return x.value.IsInt()
}

func (x Rat) Add(y Rat) Rat {
	return Rat{x.value.Add(y.value)}
}

func (x Rat) Mul(y Rat) Rat {
	return Rat{x.value.Mul(y.value)}
}

func (x Rat) RoundDown() Rat {
	return Rat{x.value.RoundDown()}
}

func (x Rat) RoundUp() Rat {
	return Rat{x.value.RoundUp()}
}

func SimplifyGameRatOptions(children collection.Set[Rat], maximize bool) collection.Set[Rat] {
	if children.IsEmpty() {
		return children
	}
	minGame := children.GetAnyElement()
	maxGame := children.GetAnyElement()
	for child := range children {
		if maxGame.Less(child) {
			maxGame = child
		}
		if minGame.Greater(child) {
			minGame = child
		}
	}
	retSet := collection.MakeSet[Rat]()
	if maximize {
		retSet.Add(maxGame)
	} else {
		retSet.Add(minGame)
	}
	return retSet
}

func DeduceGameRat(left collection.Set[Rat], right collection.Set[Rat]) Rat {
	leftSimplified := SimplifyGameRatOptions(left, true)
	rightSimplified := SimplifyGameRatOptions(right, false)
	// If no player has a move, the game is over
	if (leftSimplified.Len() == 0) && (rightSimplified.Len() == 0) {
		return MakeGameRat("0")
	}
	// If left has no move, right has the option
	// to ignore the game (=0)
	// or take one turn advantage over its best option
	if leftSimplified.Len() == 0 {
		worstRightUsableGame := MakeGameRat("1")
		rightSimplified.Add(worstRightUsableGame)
		rightSimplified = SimplifyGameRatOptions(rightSimplified, false)
		bestRightGame := rightSimplified.GetOnlyElement()
		advantage := MakeGameRat("-1")
		return bestRightGame.RoundUp().Add(advantage)
	}
	// If right has no move, left has the option
	// to ignore the game (=0)
	// or take one turn advantage over its best option
	if rightSimplified.Len() == 0 {
		worstLeftUsableGame := MakeGameRat("-1")
		leftSimplified.Add(worstLeftUsableGame)
		leftSimplified = SimplifyGameRatOptions(leftSimplified, true)
		bestLeftGame := leftSimplified.GetOnlyElement()
		advantage := MakeGameRat("1")
		return bestLeftGame.RoundDown().Add(advantage)
	}
	bestLeftGame := leftSimplified.GetOnlyElement()
	bestRightGame := rightSimplified.GetOnlyElement()
	zero := MakeGameRat("0")
	// If neither player has a cost of making the first move, we are in non-rational territory
	if !bestLeftGame.Less(bestRightGame) {
		panic("Both players want to move! This game cannot be described by a number.")
	}
	// If both players' move advance the strategy of their opponent, neither wants to move (=0).
	if (bestLeftGame.Less(zero)) && (bestRightGame.Greater(zero)) {
		return MakeGameRat("0")
	}
	// If the game is non-negative, right should only move, if it is not punished for the move advantage
	if !bestLeftGame.Less(zero) {
		disadvantage := MakeGameRat("2")
		rightSimplified.Add(bestLeftGame.RoundDown().Add(disadvantage))
		rightSimplified = SimplifyGameRatOptions(rightSimplified, false)
	}
	// If the game is non-positive, left should only move, if it is not punished for the move advantage
	if !bestRightGame.Greater(zero) {
		disadvantage := MakeGameRat("-2")
		leftSimplified.Add(bestRightGame.RoundUp().Add(disadvantage))
		leftSimplified = SimplifyGameRatOptions(leftSimplified, true)
	}
	bestLeftGame = leftSimplified.GetOnlyElement()
	bestRightGame = rightSimplified.GetOnlyElement()
	dyadicMean := bestLeftGame.value.DyadicMean(bestRightGame.value)
	return Rat{dyadicMean}
}
