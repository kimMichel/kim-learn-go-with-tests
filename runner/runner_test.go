package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {

	t.Run("compare the velocities of servers and return the address of the fastest one", func(t *testing.T) {
		slowlyServer := createDelayedServer(20 * time.Millisecond)
		fastestServer := createDelayedServer(0 * time.Millisecond)

		defer slowlyServer.Close()
		defer fastestServer.Close()

		SlowlyURL := slowlyServer.URL
		FastURL := fastestServer.URL

		expect := FastURL
		result, err := Runner(SlowlyURL, FastURL)

		if err != nil {
			t.Fatalf("doesn't expect an error, but got one %v", err)
		}

		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}

	})

	t.Run("return an error message if the server fails to respond within 10 seconds.", func(t *testing.T) {
		server := createDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := Configurable(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error, but doesn't got one")
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
