package game

import "testing"

func TestRat(t *testing.T) {
	x := MakeGameRat("7/3")
	stringActual := x.String()
	stringExpected := "7/3"
	if stringActual != stringExpected {
		t.Errorf("Stringer was wrong, got = %s; want %s", stringActual, stringExpected)
	}
}

func TestRat_Less(t *testing.T) {
	small := MakeGameRat("355/113")
	large := MakeGameRat("31416/10000")

	if !small.Less(large) {
		t.Errorf("Less was incorrect. %s should be less than %s", small, large)
	}
	if small.Less(small) {
		t.Errorf("Less was incorrect. %s should not be less than %s", small, small)
	}
	if large.Less(small) {
		t.Errorf("Less was incorrect. %s should not be less than %s", large, small)
	}
}

func TestRat_Greater(t *testing.T) {
	small := MakeGameRat("355/113")
	large := MakeGameRat("31416/10000")

	if small.Greater(large) {
		t.Errorf("Greater was incorrect. %s should not be greater than %s", small, large)
	}
	if small.Greater(small) {
		t.Errorf("Greater was incorrect. %s should not be greater than %s", small, small)
	}
	if !large.Greater(small) {
		t.Errorf("Greater was incorrect. %s should be greater than %s", large, small)
	}
}

func TestRat_IsInt(t *testing.T) {
	fracVal := MakeGameRat("999/1000")
	intVal := MakeGameRat("1000/1000")

	if fracVal.IsInt() {
		t.Errorf("IsInt was incorrect, got: %t for %s, want: %t.", fracVal.IsInt(), fracVal, false)
	}

	if !intVal.IsInt() {
		t.Errorf("IsInt was incorrect, got: %t for %s, want: %t.", intVal.IsInt(), intVal, true)
	}
}

func TestRat_Add(t *testing.T) {
	x := MakeGameRat("8/3")
	y := MakeGameRat("9/4")
	sumActual := x.Add(y)
	sumExpected := MakeGameRat("59/12")
	if sumActual != sumExpected {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", sumActual, sumExpected)
	}
}

func TestRat_Mul(t *testing.T) {
	x := MakeGameRat("8/3")
	y := MakeGameRat("7/5")
	prodActual := x.Mul(y)
	prodExpected := MakeGameRat("56/15")
	if prodActual != prodExpected {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", prodActual, prodExpected)
	}
}

func TestRat_RoundDown(t *testing.T) {
	origin := MakeGameRat("355/113")
	roundExpected := MakeGameRat("3")
	roundActual := origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeGameRat("-355/113")
	roundExpected = MakeGameRat("-4")
	roundActual = origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeGameRat("12/3")
	roundExpected = MakeGameRat("4")
	roundActual = origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
}

func TestRat_RoundUp(t *testing.T) {
	origin := MakeGameRat("355/113")
	roundExpected := MakeGameRat("4")
	roundActual := origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeGameRat("-355/113")
	roundExpected = MakeGameRat("-3")
	roundActual = origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeGameRat("12/3")
	roundExpected = MakeGameRat("4")
	roundActual = origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
}
