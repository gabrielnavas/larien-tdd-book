package array_slices_test

import (
	"larien-tdd-book/array_slices"
	"testing"
)

func TestArray(t *testing.T) {
	numbers := [5]int{1, 1, 2, 2, 3}
	got := array_slices.Sum(numbers)
	want := 9

	if got != want {
		t.Errorf("want '%d' got '%d', given '%v'", want, got, numbers)
	}
}
