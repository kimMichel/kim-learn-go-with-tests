package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repettions := Repeat("a", 6)
	expect := "aaaaaa"

	if repettions != expect {
		t.Errorf("repetitions '%s', expect '%s'", repettions, expect)
	}
}

func ExampleRepeat() {
	repetitions := Repeat("b", 6)
	fmt.Println(repetitions)
	//Output: bbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
