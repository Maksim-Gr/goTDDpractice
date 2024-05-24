package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("name", "")
		want := "Hello, name"
		asserCorrectMessageFormat(t, got, want)
	})

	t.Run("say hello default when no arg is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, default"
		asserCorrectMessageFormat(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Gopher", "Spanish")
		want := "Hola, Gopher"
		asserCorrectMessageFormat(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Gopher", "French")
		want := "Bonjour, Gopher"
		asserCorrectMessageFormat(t, got, want)
	})

}

func asserCorrectMessageFormat(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
