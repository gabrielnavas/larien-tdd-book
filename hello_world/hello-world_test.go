package helloworld_test

import (
	helloworld "larien-tdd-book/hello_world"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	verifyAllCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := helloworld.HelloWorld("John")
		want := "Olá, John"

		verifyAllCorrectMessage(t, got, want)
	})

	t.Run("Hello world, when input an empty string", func(t *testing.T) {
		got := helloworld.HelloWorld("")
		want := "Hello World"

		verifyAllCorrectMessage(t, got, want)
	})
}
