package iteration_test

import (
	"fmt"
	"larien-tdd-book/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := iteration.Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got '%s'; want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 10)
	}
}

func ExampleRepeat() {
	chars := iteration.Repeat("a", 10)
	fmt.Println(chars)
	// Output: aaaaaaaaaa
}
