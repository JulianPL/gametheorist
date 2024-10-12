package number

import "testing"

func TestNimber_Stringer(t *testing.T) {
	x := MakeNimber("123")
	stringActual := x.String()
	stringExpected := "123"
	if stringActual != stringExpected {
		t.Errorf("Stringer was wrong, got = %s; want %s", stringActual, stringExpected)
	}
}

func TestNimber_Add(t *testing.T) {
	x := MakeNimber("1")
	y := MakeNimber("7")
	stringActual := x.Add(y)
	stringExpected := MakeNimber("6")
	if stringActual != stringExpected {
		t.Errorf("Sum was incorrect, got = %s; want %s", stringActual, stringExpected)
	}
}
