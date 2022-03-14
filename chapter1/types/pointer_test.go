package types

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var a int = 100
	var b uint8 = 100
	var c float64 = 100
	var d float32 = 100
	var e string = "abcdedf"

	var pointer = &e
	var pp = &pointer

	var p3 = new(int)

	fmt.Printf("Variable pointer type=%T,address=%x\n", &a, &a)
	fmt.Printf("Variable pointer type=%T,address=%x\n", &b, &b)
	fmt.Printf("Variable pointer type=%T,address=%x\n", &c, &c)
	fmt.Printf("Variable pointer type=%T,address=%x\n", &d, &d)
	fmt.Printf("Variable pointer type=%T,address=%x\n", &e, &e)
	fmt.Printf("Variable pointer's type=%T,address=%x\n", &pointer, &pointer)
	fmt.Printf("Variable pointer's type=%T,address=%x\n", &pp, &pp)
	fmt.Printf("Variable pointer's type=%T,address=%x\n", p3, p3)
}
