package intergers_test

import (
	intergers "larien-tdd-book/integers"
	"testing"
)

func TestSum(t *testing.T) {
	got := intergers.Sum(2, 2)
	want := 4

	if got != want {
		t.Errorf("got '%d'; want '%d'", got, want)
	}
}
