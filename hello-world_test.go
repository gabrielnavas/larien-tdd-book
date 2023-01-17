package main

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Run("say hello to people", func(t *testing.T) {
		got := HelloWorld("John")
		want := "Olá, John"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
