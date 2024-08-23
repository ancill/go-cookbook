package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Ivan", "")
		want := "Hello, Ivan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ivan", "Spanish")
		want := "Hola, Ivan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Eliot", "French")
		want := "Bonjour, Eliot"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
t.Helper() is needed to tell the test suite that this method is a helper.
By doing this, when it fails, the line number reported will be in our function call
rather than inside our test helper.
*/
