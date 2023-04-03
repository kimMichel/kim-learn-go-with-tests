package runner

import (
	"net/http"
)

func Runner(a, b string) (winner string) {

	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
