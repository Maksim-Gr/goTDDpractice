package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertStringValues(t, got, want)
	})

	t.Run("search unknown word ", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "word not found"
		if err == nil {
			t.Errorf("Search should have returned err")
		}
		AssertStringValues(t, err.Error(), want)
	})
}

func AssertStringValues(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
