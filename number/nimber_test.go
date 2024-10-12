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
	sumActual := x.Add(y)
	sumExpected := MakeNimber("6")
	if sumActual != sumExpected {
		t.Errorf("Sum was incorrect, got = %s; want %s", sumActual, sumExpected)
	}
}

func TestNimber_Increment(t *testing.T) {
	x := MakeNimber("4")
	incrementActual := x.Increment()
	incrementExpected := MakeNimber("5")
	if incrementActual != incrementExpected {
		t.Errorf("Increment was incorrect, got = %s; want %s", incrementActual, incrementExpected)
	}
}
