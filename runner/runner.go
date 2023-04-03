package runner

import (
	"fmt"
	"net/http"
	"time"
)

var limitOf10Seconds = 10 * time.Second

func Runner(a, b string) (winner string, err error) {
	return Configurable(a, b, limitOf10Seconds)
}

func Configurable(a, b string, timeLimit time.Duration) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeLimit):
		return "", fmt.Errorf("timeout exceeded for %s and %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}
