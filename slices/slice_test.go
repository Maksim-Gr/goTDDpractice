package slices

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Sum(numbers)
	want := 55
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
