package array_slices_test

import (
	"larien-tdd-book/array_slices"
	"testing"
)

func TestArray(t *testing.T) {

	t.Run("collection with anything length", func(t *testing.T) {
		numbers := []int{1, 1, 2}
		got := array_slices.Sum(numbers)
		want := 4

		if got != want {
			t.Errorf("want '%d' got '%d', given '%v'", want, got, numbers)
		}
	})
}
