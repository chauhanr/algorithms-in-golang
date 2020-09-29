package hello

import "testing"

func TestHelloFunc(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello()= %q, want %q", got, want)
	}
}

func TestProverbFunc(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb()= %q, want %q", got, want)
	}
}
