package high_order

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	plus3Func := foo(3)

	plus5Func := foo(5)

	fmt.Println(plus3Func(100))
	fmt.Println(plus5Func(100))
}

func foo(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
