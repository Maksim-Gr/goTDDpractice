package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
