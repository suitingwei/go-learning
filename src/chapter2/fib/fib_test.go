package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T){
	var a int  = 1
	var b int =  1

	//fmt.Println("Start calculating....")
	t.Log(" ",a)
	for i:=0;i<=5;i++{
		//fmt.Print(" ", b)
		t.Log(" ",b)
		temp:=a
		a=b
		b=temp+a
	}
	fmt.Println()
}

func TestExchange(t *testing.T){
	a:=1
	b:=2
	//交换变量
	a,b= b,a
	t.Log(a,b)
}
