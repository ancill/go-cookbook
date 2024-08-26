package iteration

import (
	"fmt"
	"testing"
)

// to run this Benchmark go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// to run docs godoc -http=:6060
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	fmt.Println("repeated", repeated == expected)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
