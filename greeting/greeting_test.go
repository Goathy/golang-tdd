package main

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("saying ello to people", func(t *testing.T) {
		got := Greeting("Jakub")
		want := "Hello, Jakub"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' wen an empty string is supplied", func(t *testing.T) {
		got := Greeting("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
