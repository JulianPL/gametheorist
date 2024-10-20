package number

import "testing"

func TestBigInt_Stringer(t *testing.T) {
	x := MakeBigInt("7")
	stringActual := x.String()
	stringExpected := "7"
	if stringActual != stringExpected {
		t.Errorf("Stringer was wrong, got = %s; want %s", stringActual, stringExpected)
	}
}

func TestBigInt_Add(t *testing.T) {
	x := MakeBigInt("8")
	y := MakeBigInt("9")
	sumActual := x.Add(y)
	sumExpected := MakeBigInt("17")
	if sumActual != sumExpected {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", sumActual, sumExpected)
	}
}

func TestBigInt_Sub(t *testing.T) {
	x := MakeBigInt("8")
	y := MakeBigInt("9")
	sumActual := x.Sub(y)
	sumExpected := MakeBigInt("-1")
	if sumActual != sumExpected {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", sumActual, sumExpected)
	}
}

func TestBigInt_Less(t *testing.T) {
	small := MakeBigInt("16")
	large := MakeBigInt("18")

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

func TestBigInt_Greatest(t *testing.T) {
	small := MakeBigInt("16")
	large := MakeBigInt("31")

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
