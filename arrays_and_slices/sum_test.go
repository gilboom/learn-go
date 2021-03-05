package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d but want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}

	got := SumAll(slice1, slice2)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v, given %v and %v", got, want, slice1, slice2)
	}
}

func TestSumAllTail(t *testing.T) {

	checkSum := func(t testing.TB, got []int, want []int, input [][]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v, given %v, and %v", got, want, input[0], input[1])
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		got := SumAllTail(slice1, slice2)
		want := []int{2, 9}
		checkSum(t, got, want, [][]int{slice1, slice2})
	})

	t.Run("make the sums of empty slices", func(t *testing.T) {
		var slice1 []int
		slice2 := []int{3, 4, 5}

		got := SumAllTail(slice1, slice2)
		want := []int{0, 9}
		checkSum(t, got, want, [][]int{slice1, slice2})
	})
}
