package game

import (
	"gametheorist/collection"
	"gametheorist/number"
	"testing"
)

func Test_Nim(t *testing.T) {
	game := collection.MakeSet[number.Nimber]()
	game.Add(number.MakeNimber("3"))
	game.Add(number.MakeNimber("4"))
	game.Add(number.MakeNimber("6"))
	sumActual := Nim(game)
	sumExpected := MakeGameImpartial("1")
	if sumActual != sumExpected {
		t.Errorf("NimValue was incorrect, got: %v, want: %v.", sumActual, sumExpected)
	}
	game.Add(number.MakeNimber("1"))
	sumActual = Nim(game)
	sumExpected = MakeGameImpartial("0")
	if sumActual != sumExpected {
		t.Errorf("NimValue was incorrect, got: %v, want: %v.", sumActual, sumExpected)
	}
}
