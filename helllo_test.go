package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("someName")
	want := "Hello someName"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
