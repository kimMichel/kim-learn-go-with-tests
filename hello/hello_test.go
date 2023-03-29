package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	expect := "Hello World!"

	if result != expect {
		t.Errorf("result '%s', expected '%s'", result, expect)
	}
}

func TestHelloWithName(t *testing.T) {
	result := HelloWithName("Kim")
	expect := "Hello Kim"

	if result != expect {
		t.Errorf("result '%s', expected '%s'", result, expect)
	}
}
