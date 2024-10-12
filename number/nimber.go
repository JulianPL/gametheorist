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

// Nimber is a constant and comparable data structure representing arbitrary nimbers.
type Nimber struct {
	value string
}

// MakeNimber returns a Nimber with a value given by a string.
func MakeNimber(s string) Nimber {
	return Nimber{makeBigIntNative(s).String()}
}

// String returns the value of a Nimber as string.
func (x Nimber) String() string {
	return fmt.Sprintf(x.value)
}

// Add returns the sum of two Nimbers.
func (x Nimber) Add(y Nimber) Nimber {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative(y.value)
	zRaw := &big.Int{}
	zRaw.Xor(xRaw, yRaw)
	return Nimber{zRaw.String()}
}

func (x Nimber) Increment() Nimber {
	xRaw := makeBigIntNative(x.value)
	yRaw := makeBigIntNative("1")
	zRaw := &big.Int{}
	zRaw.Add(xRaw, yRaw)
	return Nimber{zRaw.String()}
}
