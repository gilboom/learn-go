package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := new(bytes.Buffer)
	Greet(buffer, "GilBoom")

	got := buffer.String()
	want := "Hello, GilBoom"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
