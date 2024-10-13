package game

import (
	"gametheorist/collection"
	"gametheorist/number"
)

func NimValue(stacks collection.Set[number.Nimber]) number.Nimber {
	value := number.MakeNimber("0")
	for stack := range stacks {
		value = value.Add(stack)
	}
	return value
}
