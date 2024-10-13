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
