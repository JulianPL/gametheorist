package number

import "testing"

func TestBigRat_Stringer(t *testing.T) {
	x := MakeBigRat("7/3")
	stringActual := x.String()
	stringExpected := "7/3"
	if stringActual != stringExpected {
		t.Errorf("Stringer was wrong, got = %s; want %s", stringActual, stringExpected)
	}
}

func TestBigRat_MakeBigRat(t *testing.T) {
	x := MakeBigRat("8/4")
	y := MakeBigRat("2/1")

	if x != y {
		t.Errorf("Creation should lead to canonical representation, but: %s != %s detected.", x, y)
	}
}

func TestBigRat_Add(t *testing.T) {
	x := MakeBigRat("8/3")
	y := MakeBigRat("9/4")
	sumActual := x.Add(y)
	sumExpected := MakeBigRat("59/12")
	if sumActual != sumExpected {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", sumActual, sumExpected)
	}
}

func TestBigRat_Sub(t *testing.T) {
	x := MakeBigRat("8/3")
	y := MakeBigRat("9/4")
	diffActual := x.Sub(y)
	diffExpected := MakeBigRat("5/12")
	if diffActual != diffExpected {
		t.Errorf("Diff was incorrect, got: %s, want: %s.", diffActual, diffExpected)
	}
}

func TestBigRat_Mul(t *testing.T) {
	x := MakeBigRat("8/3")
	y := MakeBigRat("7/5")
	prodActual := x.Mul(y)
	prodExpected := MakeBigRat("56/15")
	if prodActual != prodExpected {
		t.Errorf("Mul was incorrect, got: %s, want: %s.", prodActual, prodExpected)
	}
}

func TestBigRat_IsInt(t *testing.T) {
	fracVal := MakeBigRat("999/1000")
	intVal := MakeBigRat("1000/1000")

	if fracVal.IsInt() {
		t.Errorf("IsInt was incorrect, got: %t for %s, want: %t.", fracVal.IsInt(), fracVal, false)
	}

	if !intVal.IsInt() {
		t.Errorf("IsInt was incorrect, got: %t for %s, want: %t.", intVal.IsInt(), intVal, true)
	}
}

func TestBigRat_Less(t *testing.T) {
	small := MakeBigRat("355/113")
	large := MakeBigRat("31416/10000")

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

func TestBigRat_Greater(t *testing.T) {
	small := MakeBigRat("355/113")
	large := MakeBigRat("31416/10000")

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

func TestBigRat_RoundDown(t *testing.T) {
	origin := MakeBigRat("355/113")
	roundExpected := MakeBigRat("3")
	roundActual := origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeBigRat("-355/113")
	roundExpected = MakeBigRat("-4")
	roundActual = origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeBigRat("12/3")
	roundExpected = MakeBigRat("4")
	roundActual = origin.RoundDown()
	if roundActual != roundExpected {
		t.Errorf("RoundDown was incorrect. %s != %s", roundActual, roundExpected)
	}
}

func TestBigRat_RoundUp(t *testing.T) {
	origin := MakeBigRat("355/113")
	roundExpected := MakeBigRat("4")
	roundActual := origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeBigRat("-355/113")
	roundExpected = MakeBigRat("-3")
	roundActual = origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
	origin = MakeBigRat("12/3")
	roundExpected = MakeBigRat("4")
	roundActual = origin.RoundUp()
	if roundActual != roundExpected {
		t.Errorf("RoundUp was incorrect. %s != %s", roundActual, roundExpected)
	}
}
