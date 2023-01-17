package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("John")
	want := "Hello John"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
