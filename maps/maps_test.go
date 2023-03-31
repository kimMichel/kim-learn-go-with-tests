package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("found word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expect := "this is just a test"

		compareString(t, result, expect)
	})

	t.Run("not found word", func(t *testing.T) {
		_, result := dictionary.Search("unknown")

		compareError(t, result, ErrNotFound)
	})
}

func compareString(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("result '%s', expect '%s', passed '%s'", result, expect, "test")
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("result error '%s', expected '%s'", result, expect)
	}
}
