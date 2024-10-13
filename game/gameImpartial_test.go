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

func Test_TheWhiteKnight(t *testing.T) {
	left := 1
	right := 2
	gameActual := TheWhiteKnight(left, right)
	gameExpected := MakeGameImpartial("0")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}

	left = 5
	right = 4
	gameActual = TheWhiteKnight(left, right)
	gameExpected = MakeGameImpartial("4")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}

	left = 5
	right = 3
	gameActual = TheWhiteKnight(left, right)
	gameExpected = MakeGameImpartial("3")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}

	left = 9
	right = 10
	gameActual = TheWhiteKnight(left, right)
	gameExpected = MakeGameImpartial("0")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}

	left = 14
	right = 7
	gameActual = TheWhiteKnight(left, right)
	gameExpected = MakeGameImpartial("2")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}

	left = 15
	right = 9
	gameActual = TheWhiteKnight(left, right)
	gameExpected = MakeGameImpartial("1")
	if gameActual != gameExpected {
		t.Errorf("TheWhiteKnight was incorrect, got: %v, want: %v.", gameActual, gameExpected)
	}
}
