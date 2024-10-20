package collection

import (
	"reflect"
	"testing"
)

func TestMakeOrderedPartitions(t *testing.T) {
	actual := len(MakeOrderedPartitions(7, 3))
	expected := 15
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MakeOrderedPartitions was wrong, got %v, want %v", actual, expected)
	}

	actual = len(MakeOrderedPartitions(7, 8))
	expected = 0
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MakeOrderedPartitions was wrong, got %v, want %v", actual, expected)
	}

	actual = len(MakeOrderedPartitions(0, 0))
	expected = 1
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MakeOrderedPartitions was wrong, got %v, want %v", actual, expected)
	}

	actual = len(MakeOrderedPartitions(7, 0))
	expected = 0
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MakeOrderedPartitions was wrong, got %v, want %v", actual, expected)
	}
}
