package main

import "testing"

func TestHelloWorld(t *testing.T) {

	verifyAllCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := HelloWorld("John")
		want := "Olá, John"

		verifyAllCorrectMessage(t, got, want)
	})

	t.Run("Hello world, when input an empty string", func(t *testing.T) {
		got := HelloWorld("")
		want := "Hello World"

		verifyAllCorrectMessage(t, got, want)
	})
}
