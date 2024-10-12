package game

import (
	"gametheorist/collection"
	"strconv"
)

func CutCake(left, right int) Rat {
	a := make([][]Rat, left+1)
	for i := range a {
		a[i] = make([]Rat, right+1)
	}
	for l := 0; l <= left; l++ {
		for r := 0; r <= right; r++ {
			if l == 0 || r == 0 {
				a[l][r] = MakeGameRat("0")
				continue
			}
			if l == 1 && r == 1 {
				a[l][r] = MakeGameRat("0")
				continue
			}
			if l == 1 {
				a[l][r] = a[l][r-1].Add(MakeGameRat("-1"))
				continue
			}
			if r == 1 {
				a[l][r] = a[l-1][r].Add(MakeGameRat("1"))
				continue
			}
			leftOptions := collection.MakeSet[Rat]()
			rightOptions := collection.MakeSet[Rat]()
			for lCurr := 1; lCurr < l; lCurr++ {
				leftOptions.Add(a[lCurr][r].Add(a[l-lCurr][r]))
			}
			for rCurr := 1; rCurr < r; rCurr++ {
				rightOptions.Add(a[l][rCurr].Add(a[l][r-rCurr]))
			}
			a[l][r] = DeduceGameRat(leftOptions, rightOptions)
		}
	}
	return a[left][right]
}

func MaundyCake(left, right int) Rat {
	a := make([][]Rat, left+1)
	for i := range a {
		a[i] = make([]Rat, right+1)
	}
	for l := 0; l <= left; l++ {
		for r := 0; r <= right; r++ {
			if l == 0 || r == 0 {
				a[l][r] = MakeGameRat("0")
				continue
			}
			if l == 1 && r == 1 {
				a[l][r] = MakeGameRat("0")
				continue
			}
			if l == 1 {
				a[l][r] = a[l][r-1].Add(MakeGameRat("-1"))
				continue
			}
			if r == 1 {
				a[l][r] = a[l-1][r].Add(MakeGameRat("1"))
				continue
			}
			leftOptions := collection.MakeSet[Rat]()
			rightOptions := collection.MakeSet[Rat]()
			for lCurr := 1; lCurr < l; lCurr++ {
				if l%lCurr == 0 {
					leftOptions.Add(a[lCurr][r].Mul(MakeGameRat(strconv.Itoa(l / lCurr))))
				}
			}
			for rCurr := 1; rCurr < r; rCurr++ {
				if r%rCurr == 0 {
					rightOptions.Add(a[l][rCurr].Mul(MakeGameRat(strconv.Itoa(r / rCurr))))
				}
			}
			a[l][r] = DeduceGameRat(leftOptions, rightOptions)
		}
	}
	return a[left][right]
}
