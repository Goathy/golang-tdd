package generics

import "testing"

func AssertEqual[Type comparable](t *testing.T, got, want Type) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[Type comparable](t *testing.T, got, want Type) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
