package game

import "testing"

func Test_CutCake(t *testing.T) {
	left := 0
	right := 0
	gameActual := CutCake(left, right)
	gameExpected := MakeGameRat("0")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 1
	right = 1
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("0")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 13
	right = 1
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("12")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 1
	right = 4
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("-3")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 7
	right = 4
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("0")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 8
	right = 4
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("1")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 3
	right = 13
	gameActual = CutCake(left, right)
	gameExpected = MakeGameRat("-5")
	if gameActual != gameExpected {
		t.Errorf("CutCake was incorrect, got: %v=CutCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}
}

func Test_MaundyCake(t *testing.T) {
	left := 16
	right := 1
	gameActual := MaundyCake(left, right)
	gameExpected := MakeGameRat("15")
	if gameActual != gameExpected {
		t.Errorf("MaundyCake was incorrect, got: %v=MaundyCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 14
	right = 5
	gameActual = MaundyCake(left, right)
	gameExpected = MakeGameRat("1")
	if gameActual != gameExpected {
		t.Errorf("MaundyCake was incorrect, got: %v=MaundyCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 13
	right = 7
	gameActual = MaundyCake(left, right)
	gameExpected = MakeGameRat("0")
	if gameActual != gameExpected {
		t.Errorf("MaundyCake was incorrect, got: %v=MaundyCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 11
	right = 16
	gameActual = MaundyCake(left, right)
	gameExpected = MakeGameRat("-7")
	if gameActual != gameExpected {
		t.Errorf("MaundyCake was incorrect, got: %v=MaundyCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}

	left = 16
	right = 9
	gameActual = MaundyCake(left, right)
	gameExpected = MakeGameRat("3")
	if gameActual != gameExpected {
		t.Errorf("MaundyCake was incorrect, got: %v=MaundyCake(%d, %d), want: %v.", gameActual, left, right, gameExpected)
	}
}
