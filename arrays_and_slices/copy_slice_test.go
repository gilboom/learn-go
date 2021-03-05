package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestCopyStrings(t *testing.T) {
	slice := []string{"GilBoom", "Phantom"}
	got := CopyStrings(slice)
	want := []string{"GilBoom", "Phantom"}
	t.Logf("original slice potiner %p, new slice pointer %p", &slice, &got)
	if &slice == &got {
		t.Errorf("expect original slice pointer not equals to new slice pointer, but equal")
	}
	if !reflect.DeepEqual(slice, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}
