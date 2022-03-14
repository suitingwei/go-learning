package structs

import (
	"context"
	"fmt"
	"testing"
)

type Base struct {
	ctx      context.Context
	baseName string
}

//func (receiver Base) BaseName()  string  {
//	return receiver.baseName
//}

type A struct {
	*Base
}

func (receiver A) aFunc() {
	fmt.Println("this is function A called in a struct, i can reach the base field,", receiver.baseName)
}

type B struct {
	*Base
}

func (receiver B) bFunc() {
	fmt.Println("this is function B called in B struct, i can reach the base field,", receiver.baseName)
}

type Client struct {
	*Base
	*A
	*B
}

func TestName(t *testing.T) {
	baseInstance := &Base{ctx: context.Background(), baseName: "i am the base name of the shared object"}
	client := &Client{
		Base: baseInstance,
		A:    &A{},
		B:    &B{},
	}

	client.aFunc()
	client.bFunc()
}

func TestStrictType(t *testing.T) {
	type AgeRange int

	var young AgeRange = 100
	var old AgeRange = 999
	t.Log(young, old)

	intPrinter := func(age int) {
		fmt.Printf("Age is %d\n", age)
	}
	intPrinter(int(young))
	intPrinter(int(old))
}
