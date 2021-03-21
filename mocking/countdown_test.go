package main

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := CountdownOperationsSpy{}

		Countdown(buffer, &spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := CountdownOperationsSpy{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		got := spySleepPrinter.Calls
		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v but got %v", want, got)
		}
	})
}
