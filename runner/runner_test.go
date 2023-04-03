package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {

	slowlyServer := createDelayedServer(20 * time.Millisecond)
	fastestServer := createDelayedServer(0 * time.Millisecond)

	// Ao chamar uma função com o prefixo defer, ela será chamada após o termino da função que a contém.
	defer slowlyServer.Close()
	defer fastestServer.Close()

	SlowlyURL := slowlyServer.URL
	FastURL := fastestServer.URL

	expect := FastURL
	result := Runner(SlowlyURL, FastURL)

	if result != expect {
		t.Errorf("result '%s', expected '%s'", result, expect)
	}
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
