package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("name")
		want := "Hello, name"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("say hello default when no arg is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, default"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
