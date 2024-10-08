package collection

import "testing"

func TestSet_Contains(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	if !testSet.Contains(1) {
		t.Errorf("Add or Contains was incorrect, element %d not detected", 1)
	}
	if testSet.Contains(2) {
		t.Errorf("Contains was incorrect, element %d detected", 2)
	}
}

func TestSet_Add(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	if !testSet.Contains(1) {
		t.Errorf("Add or Contains was incorrect, element %d not detected", 1)
	}
}

func TestSet_Remove(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	testSet.Remove(1)
	if testSet.Contains(1) {
		t.Errorf("Remove or Contains was incorrect, element %d detected", 1)
	}
}

func TestSet_Clear(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	testSet.Clear()
	if testSet.Contains(1) {
		t.Errorf("Remove or Contains was incorrect, element %d detected", 1)
	}
}

func TestSet_Len(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	testSet.Add(2)
	testSet.Add(3)
	testSet.Add(4)
	testSet.Remove(3)
	lenActual := testSet.Len()
	lenExpected := 3

	if lenActual != lenExpected {
		t.Errorf("Length was incorrect, got: %d, want: %d.", lenActual, lenExpected)
	}
}

func TestSet_IsEmpty(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	if testSet.IsEmpty() {
		t.Errorf("IsEmpty was incorrect, set with one element reported empty")
	}
	testSet.Remove(1)
	if !testSet.IsEmpty() {
		t.Errorf("IsEmpty was incorrect, set without elements reported non-empty")
	}
}

func TestSet_GetAnyElement(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	testSet.Add(2)
	testSet.Add(3)
	element := testSet.GetAnyElement()
	if !testSet.Contains(element) {
		t.Errorf("GetAnyElement was incorrect, given element %d not detected", element)
	}
}

func TestSet_GetAnyElement_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAnyElement did not panic on empty List")
		}
	}()

	testSet := MakeSet[int]()
	testSet.GetAnyElement()
}

func TestSet_GetOneElement(t *testing.T) {
	testSet := MakeSet[int]()
	testSet.Add(1)
	element := testSet.GetOnlyElement()
	if !testSet.Contains(element) {
		t.Errorf("GetOneElement was incorrect, given element %d not detected", element)
	}
}

func TestSet_GetOneElement_PanicEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetOnlyElement did not panic on empty List")
		}
	}()

	testSet := MakeSet[int]()
	testSet.GetOnlyElement()
}

func TestSet_GetOneElement_PanicMultiple(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetOnlyElement did not panic on List with more than one element")
		}
	}()

	testSet := MakeSet[int]()
	testSet.Add(1)
	testSet.Add(2)
	testSet.GetOnlyElement()
}
