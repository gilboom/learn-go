package test_based_property

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Name   string
		Arabic int
		Want   string
	}{
		{"1 get converted to I", 1, "I"},
		{"2 get converted to II", 2, "II"},
		{"3 get converted to III", 3, "III"},
		{"4 get converted to IV", 4, "IV"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := ConvertToRoman(c.Arabic)

			if got != c.Want {
				t.Errorf("got %q but want %q", got, c.Want)
			}
		})
	}

	//t.Run("1 get converted to I", func(t *testing.T) {
	//	got := ConvertToRoman(1)
	//	want := "I"
	//
	//	if got != want {
	//		t.Errorf("got %q but want %q", got, want)
	//	}
	//})
	//
	//t.Run("2 get converted to II", func(t *testing.T) {
	//	got := ConvertToRoman(2)
	//	want := "II"
	//
	//	if got != want {
	//		t.Errorf("got %q but want %q", got, want)
	//	}
	//})
}
