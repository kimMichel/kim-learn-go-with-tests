package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repettions := Repeat("a")
	expect := "aaaaa"

	if repettions != expect {
		t.Errorf("repetitions '%s', expect '%s'", repettions, expect)
	}
}

func ExampleRepeat() {
	repetitions := Repeat("b")
	fmt.Println(repetitions)
	//Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
