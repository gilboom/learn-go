package arrays_and_slices

import "testing"

func TestMakeSlice(t *testing.T) {
	s := make([]int, 0, 5)
	want := 0
	got := len(s)
	if got != want {
		t.Errorf("want s length %d but got %d", want, got)
	}
}
