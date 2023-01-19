package iteration_test

import (
	"larien-tdd-book/iteration"
	"testing"
)

func TestIteration(t *testing.T) {
	got := iteration.Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got '%s'; want '%s'", got, want)
	}
}
