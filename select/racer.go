package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimout = 10 * time.Second

func Racer(a, b string) (string, error) {
	//var winner string
	//aDuration := measureResponseTime(a)
	//bDuration := measureResponseTime(b)
	//
	//if aDuration > bDuration {
	//	winner = b
	//} else {
	//	winner = a
	//}
	//return winner, nil

	return ConfigurableRacer(a, b, tenSecondTimout)
}

func GetChannel() <-chan string {
	ch := make(chan string)
	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		ch <- struct{}{}
		close(ch)
	}()
	return ch
}
