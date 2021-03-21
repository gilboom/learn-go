package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		expectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 25},
			[]string{"Chris"},
		},
		{
			"Struct with nested struct field",
			Person{
				"Chris",
				Profile{25, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{25, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{25, "Shanghai"},
			},
			[]string{"London", "Shanghai"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Shanghai"},
			},
			[]string{"London", "Shanghai"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.expectedCalls) {
				t.Errorf("got %v but want %v", got, test.expectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	//expected := "Chris"
	//var got []string
	//
	//x := struct {
	//	Name string
	//}{expected}
	//
	//Walk(x, func(input string) {
	//	got = append(got, input)
	//})
	//
	//if len(got) != 1 {
	//	t.Errorf("wrong number of function calls, got %d but want %d", len(got), 1)
	//}
	//
	//if got[0] != expected {
	//	t.Errorf("got %q but want %q", got[0], expected)
	//}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
