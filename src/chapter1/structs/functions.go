package structs

import "fmt"

//这个真的狠
//还好我特么之前对于 php 的闭包比较熟悉了。这个就是传入一个闭包，然后返回另一个闭包
func FuncWrapper(simpleEcho func(str string)) func(str string) {
	return func(str string) {
		fmt.Println("Before echo action")
		simpleEcho(str)
		fmt.Println("After echo action")
	}
}
