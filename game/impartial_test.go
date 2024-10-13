package game

import (
	"gametheorist/collection"
	"testing"
)

func TestImpartial(t *testing.T) {
	x := MakeGameImpartial("4")
	stringActual := x.String()
	stringExpected := "4"
	if stringActual != stringExpected {
		t.Errorf("Stringer was wrong, got = %s; want %s", stringActual, stringExpected)
	}
}

func TestImpartial_Increment(t *testing.T) {
	x := MakeGameImpartial("4")
	incrementActual := x.Increment()
	incrementExpected := MakeGameImpartial("5")
	if incrementActual != incrementExpected {
		t.Errorf("Increment was incorrect, got = %s; want %s", incrementActual, incrementExpected)
	}
}

func TestImpartial_MinimalExcluded(t *testing.T) {
	children := collection.MakeSet[Impartial]()
	mexActual := MinimalExcluded(children)
	mexExpected := MakeGameImpartial("0")
	if mexActual != mexExpected {
		t.Errorf("MinimalExcluded was incorrect, got = %s; want %s", mexActual, mexExpected)
	}

	children.Add(MakeGameImpartial("0"))
	children.Add(MakeGameImpartial("3"))
	children.Add(MakeGameImpartial("2"))
	children.Add(MakeGameImpartial("6"))
	children.Add(MakeGameImpartial("5"))
	children.Add(MakeGameImpartial("1"))
	mexActual = MinimalExcluded(children)
	mexExpected = MakeGameImpartial("4")
	if mexActual != mexExpected {
		t.Errorf("MinimalExcluded was incorrect, got = %s; want %s", mexActual, mexExpected)
	}
}
