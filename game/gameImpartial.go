package game

import (
	"gametheorist/collection"
	"gametheorist/number"
)

func Nim(stacks collection.Set[number.Nimber]) Impartial {
	value := MakeGameImpartial("0")
	for stack := range stacks {
		value = value.Add(Impartial{stack})
	}
	return value
}

func TheWhiteKnight(left, right int) Impartial {
	a := make([][]Impartial, left+right+2)
	for i := range a {
		a[i] = make([]Impartial, left+right+2)
	}
	for sum := 2; sum <= left+right+2; sum++ {
		for l := 1; l <= left+right+1; l++ {
			r := sum - l
			if r <= 0 {
				break
			}
			options := collection.MakeSet[Impartial]()
			if l-2 > 0 && r-1 > 0 {
				options.Add(a[l-2][r-1])
			}
			if l-2 > 0 && r+1 > 0 {
				options.Add(a[l-2][r+1])
			}
			if l-1 > 0 && r-2 > 0 {
				options.Add(a[l-1][r-2])
			}
			if l+1 > 0 && r-2 > 0 {
				options.Add(a[l+1][r-2])
			}
			a[l][r] = MinimalExcluded(options)
		}
	}
	return a[left][right]
}
