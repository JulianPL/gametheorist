package game

import (
	"gametheorist/collection"
	"testing"
)

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

func TestRat_SimplifyGameRatOptions(t *testing.T) {
	children := collection.MakeSet[Rat]()
	actualSet := SimplifyGameRatOptions(children, true)
	if !actualSet.IsEmpty() {
		t.Errorf("SimplifyGameRatOptions was incorrect, got non-empty set from empty set")
	}

	children.Add(MakeGameRat("0"))
	children.Add(MakeGameRat("-2"))
	children.Add(MakeGameRat("2"))
	children.Add(MakeGameRat("5/2"))
	children.Add(MakeGameRat("9/4"))
	actualSet = SimplifyGameRatOptions(children, true)
	actualGame := actualSet.GetOnlyElement()
	expectedGame := MakeGameRat("5/2")
	if actualGame != expectedGame {
		t.Errorf("SimplifyGameRatOptions was incorrect, got: %v, want: %v.", actualGame, expectedGame)
	}
	actualSet = SimplifyGameRatOptions(children, false)
	actualGame = actualSet.GetOnlyElement()
	expectedGame = MakeGameRat("-2")
	if actualGame != expectedGame {
		t.Errorf("SimplifyGameRatOptions was incorrect, got: %v, want: %v.", actualGame, expectedGame)
	}
}

func TestRat_DeduceGameRat_Empty(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for empty game, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_RightOptional(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("5/4"))
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional right move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("1/4"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional right move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("-1"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("-2")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional right move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("-5/4"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("-2")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional right move, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_LeftOptional(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("-5/4"))
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional left move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("-1/4"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional left move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("1"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("2")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional left move, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("5/4"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("2")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for optional left move, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_TwoSidedDisadvantage(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("-1/4"))
	rightChildren.Add(MakeGameRat("1/8"))
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("0")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for two-sided disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_RightDisadvantage(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("1/4"))
	rightChildren.Add(MakeGameRat("6"))
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("1")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for right disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("2"))
	rightChildren.Add(MakeGameRat("6"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("3")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for right disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("5/4"))
	rightChildren.Add(MakeGameRat("3/2"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("11/8")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for right disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_LeftDisadvantage(t *testing.T) {
	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("-1/4"))
	leftChildren.Add(MakeGameRat("-6"))
	GameActual := DeduceGameRat(leftChildren, rightChildren)
	GameExpected := MakeGameRat("-1")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for left disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("-2"))
	leftChildren.Add(MakeGameRat("-6"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("-3")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for left disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}

	leftChildren = collection.MakeSet[Rat]()
	rightChildren = collection.MakeSet[Rat]()
	rightChildren.Add(MakeGameRat("-5/4"))
	leftChildren.Add(MakeGameRat("-3/2"))
	GameActual = DeduceGameRat(leftChildren, rightChildren)
	GameExpected = MakeGameRat("-11/8")
	if GameActual != GameExpected {
		t.Errorf("DeduceGameRat was incorrect for left disadvantage, got: %v, want: %v.", GameActual, GameExpected)
	}
}

func TestRat_DeduceGameRat_Panic_TwoSidedAdvantage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetOnlyElement did not panic on List with more than one element")
		}
	}()

	leftChildren := collection.MakeSet[Rat]()
	rightChildren := collection.MakeSet[Rat]()
	leftChildren.Add(MakeGameRat("1/4"))
	rightChildren.Add(MakeGameRat("1/4"))
	DeduceGameRat(leftChildren, rightChildren)
}
