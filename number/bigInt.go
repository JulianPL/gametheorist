package number

import (
	"fmt"
	"math/big"
)

// makeBigIntNative returns a pointer to a big.Int with a value given by a string.
func makeBigIntNative(s string) *big.Int {
	res := &big.Int{}
	res.SetString(s, 10)
	return res
}

// BigInt is a constant and comparable data structure representing arbitrary integer numbers.
type BigInt struct {
	value string
}

// MakeBigInt returns a BigInt with a value given by a string.
func MakeBigInt(s string) BigInt {
	return BigInt{makeBigIntNative(s).String()}
}

// String returns the value of a BigInt as string.
func (x BigInt) String() string {
	return fmt.Sprintf(x.value)
}

// Add returns the sum of two BigInts.
func (x BigInt) Add(y BigInt) BigInt {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative(y.value)
	zRaw := &big.Int{}
	zRaw.Add(xRaw, yRaw)
	return BigInt{zRaw.String()}
}

func (x BigInt) Sub(y BigInt) BigInt {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative(y.value)
	zRaw := &big.Int{}
	zRaw.Sub(xRaw, yRaw)
	return BigInt{zRaw.String()}
}

func (x BigInt) Less(y BigInt) bool {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative(y.value)
	return xRaw.Cmp(yRaw) == -1
}

func (x BigInt) Greater(y BigInt) bool {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative(y.value)
	return xRaw.Cmp(yRaw) == 1
}
