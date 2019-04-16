package _const

import (
	"testing"
)

//连续常量赋值,类似 php 的数组数字键自动增长
const (
	Monday = iota+1
	Tuesday
	Wednesday
)

//const (
//	Monday = 1
//	Tuesday  = 2
//	Wednesday  =3
//)

func TestPointer(t *testing.T){

	a:=2
	aPointer:= &a

	t.Log(a,aPointer)
}
