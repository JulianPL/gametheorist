// Package number implements comparable data structures representing arbitrarily complex numbers.
package number

import (
	"fmt"
	"math/big"
)

// makeBigRatNative returns a pointer to a big.Rat with a value given by a string.
// Note that big.Rat is not the bigRat of this package nor is it comparable.
func makeBigRatNative(s string) *big.Rat {
	res := &big.Rat{}
	res.SetString(s)
	return res
}

// BigRat is a constant and comparable data structure representing arbitrary rational numbers.
type BigRat struct {
	value string
}

// MakeBigRat returns a BigRat with a value given by a string.
func MakeBigRat(s string) BigRat {
	return BigRat{makeBigRatNative(s).String()}
}

// String returns the value of a BigRat as string.
func (x BigRat) String() string {
	return fmt.Sprintf(x.value)
}

// Add returns the sum of two BigRats.
func (x BigRat) Add(y BigRat) BigRat {
	xRaw := makeBigRatNative(x.value)
	yRaw := makeBigRatNative(y.value)
	zRaw := &big.Rat{}
	zRaw.Add(xRaw, yRaw)
	return BigRat{zRaw.String()}
}

// Sub returns the difference of two BigRats.
func (x BigRat) Sub(y BigRat) BigRat {
	xRaw := makeBigRatNative(x.value)
	yRaw := makeBigRatNative(y.value)
	zRaw := &big.Rat{}
	zRaw.Sub(xRaw, yRaw)
	return BigRat{zRaw.String()}
}

// Mul returns the product of two BigRats.
func (x BigRat) Mul(y BigRat) BigRat {
	xRaw := makeBigRatNative(x.value)
	yRaw := makeBigRatNative(y.value)
	zRaw := &big.Rat{}
	zRaw.Mul(xRaw, yRaw)
	return BigRat{zRaw.String()}
}

// IsInt return true if and only if the denominator of the given BigRat is 1 in its shortened form.
func (x BigRat) IsInt() bool {
	xRaw := makeBigRatNative(x.value)
	return xRaw.IsInt()
}

func (x BigRat) Less(y BigRat) bool {
	xRaw := makeBigRatNative(x.value)
	yRaw := makeBigRatNative(y.value)
	return xRaw.Cmp(yRaw) == -1
}

func (x BigRat) Greater(y BigRat) bool {
	xRaw := makeBigRatNative(x.value)
	yRaw := makeBigRatNative(y.value)
	return xRaw.Cmp(yRaw) == 1
}

// RoundDown returns the largest integer which is not greater than the given BigRat.
func (x BigRat) RoundDown() BigRat {
	if x.IsInt() {
		return x
	}
	zero := MakeBigRat("0")
	if x.Less(zero) {
		negX := zero.Sub(x)
		return zero.Sub(negX.RoundUp())
	}
	xRaw := makeBigRatNative(x.value)
	num := xRaw.Num()
	denom := xRaw.Denom()
	mod := big.NewInt(0)
	mod.Mod(num, denom)
	num.Sub(num, mod)
	xRaw.SetFrac(num, denom)
	return MakeBigRat(xRaw.String())
}

// RoundUp returns the smallest integer which is not less than the given BigRat.
func (x BigRat) RoundUp() BigRat {
	if x.IsInt() {
		return x
	}
	one := MakeBigRat("1")
	return one.Add(x.RoundDown())
}

// DyadicMean returns the most simple dyadic number between two BigRats
// DyadicMean is only well-defined if there is at most one integer between x and y
func (x BigRat) DyadicMean(y BigRat) BigRat {
	if x == y {
		return x
	}
	if x.Greater(y) {
		return y.DyadicMean(x)
	}
	if (x.RoundDown().Add(MakeBigRat("2"))).Less(y.RoundUp()) {
		errorString := fmt.Sprintf("DyadicMean is not well-defined since is at more than one integer between %s and %s", x, y)
		panic(errorString)
	}
	if x.RoundUp() == y.RoundDown() {
		return x.RoundUp()
	}
	current := x.RoundDown()
	offset := MakeBigRat("2")
	half := MakeBigRat("1/2")
	for !current.Greater(x) {
		offset = offset.Mul(half)
		temp := current.Add(offset)
		if temp.Less(y) {
			current = temp
		}
	}
	return current
}
