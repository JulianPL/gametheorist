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
