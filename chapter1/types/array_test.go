package types

import (
	"math/rand"
	"testing"
)

const matrixLength = 2048

func getSrcArray() [matrixLength][matrixLength]int {
	result := [matrixLength][matrixLength]int{}
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			result[i][j] = rand.Int()
		}
	}
	return result
}

func BenchmarkRowCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		src := getSrcArray()
		dest := [matrixLength][matrixLength]int{}

		rowCopy(src, dest)
	}
}
func BenchmarkColumnCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		src := getSrcArray()
		dest := [matrixLength][matrixLength]int{}

		rowCopy(src, dest)
	}
}

func rowCopy(src, dest [matrixLength][matrixLength]int) {
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			dest[i][j] = src[i][j]
		}
	}
}
func columnCopy(src, dest [matrixLength][matrixLength]int) {
	for j := 0; j < matrixLength; j++ {
		for i := 0; i < matrixLength; i++ {
			dest[i][j] = src[i][j]
		}
	}
}
