package high_order

import (
	"fmt"
	"testing"
)

//简单的高阶函数，其实可以看出高阶函数的套路
//  //假如有普通函数A
//  funcA(args ...interface{}) response
//
//  //那么他对应的高阶函数，参数就是funcA、返回值也是funcA,最终就达到了装饰的效果
//  funcB(args ...interface{}) funcA
func TestName(t *testing.T) {
	plus3Func := foo(3)

	plus5Func := foo(5)

	fmt.Println(plus3Func(100))
	fmt.Println(plus5Func(100))
}

//普通的高阶函数,没有封装低阶函数，只是生成一个函数。
func foo(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//2.空字符串装饰器

//简单函数，只是输出字符串,但是不确认字符串是否为空
func printStr(s string) {
	fmt.Println(s)
}

//函数装饰器，不直接执行函数，而是返回装饰后的闭包用于调用方执行
func EmptyStrDecorator(fn func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Before decorator")
		fn(s)
		fmt.Println("After decorator")
	}
}

func TestDecorator(t *testing.T) {
	printStr("i am dog")

}
