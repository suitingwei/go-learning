package performance

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	num := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(num), 10)
	}
}
