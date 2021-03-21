package reflection

import "testing"

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d but want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q but want %q", got[0], expected)
	}
}
