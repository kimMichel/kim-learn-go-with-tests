package concurrecy

import (
	"reflect"
	"testing"
	"time"
)

func mockVerificadorWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestVerifyWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expect := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	result := VerifyWebsite(mockVerificadorWebsite, websites)

	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expected %v, result %v", expect, result)
	}
}

func slowStubVerifyWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerifyWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		VerifyWebsite(slowStubVerifyWebsite, urls)
	}
}
