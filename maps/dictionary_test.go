package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test value"}

	got := Search(dictionary, "test")
	want := "this is a test value"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
