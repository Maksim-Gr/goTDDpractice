package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test value"}

	got := Search(dictionary, "test")
	want := "this is a test value"
	AssertStringValues(t, got, want)
}

func AssertStringValues(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
