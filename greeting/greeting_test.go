package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Jakub")
	want := "Hello, Jakub"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
