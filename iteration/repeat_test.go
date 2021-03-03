package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but repeated %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	repeatStr := "0123456789"
	for _, length := range []int{5, 10} {
		for _, times := range []uint{1, 2, 6} {
			b.Run(fmt.Sprintf("string with %d length repeat %d times", length, times), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Repeat(repeatStr[:length], times)
				}
			})
		}
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 5))
	// Output: bbbbb
}
