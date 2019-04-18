package main

import "fmt"
import "./structs"

func init() {
	fmt.Println("The init function has been triggered.")
}

func main() {

	//调用了内部其他包的方法、结构体。注意这里只有首字母大写的才是public 的，而且和 private 的方法还是不一样的
	//可以同时有 public 和 private 两个名字一样的方法
	father := structs.User{"John", "John@email.com"}

	structs.LearnNestingStruct()

	father.Notify()
}

func printArray(arr *[5]int) {
	arrLen := len(*arr)

	for i := 0; i < arrLen; i++ {
		fmt.Print(" ", (*arr)[i])
	}
}
