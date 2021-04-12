package main

import "testing"

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &Tape{file: file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
}
