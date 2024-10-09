package game

import (
	"fmt"
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
