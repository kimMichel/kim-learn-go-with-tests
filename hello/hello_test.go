package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%s', expected '%s'", result, expect)
		}
	}

	t.Run("say Hello to people", func(t *testing.T) {
		result := Hello("Kim")
		expect := "Hello, Kim"

		if result != expect {
			verifyCorrectMessage(t, result, expect)
		}
	})

	t.Run("'World' has default to empty string", func(t *testing.T) {
		result := Hello("")
		expect := "Hello, World"

		if result != expect {
			verifyCorrectMessage(t, result, expect)
		}
	})
}
