package main

import (
	"fmt"
	"os"
)

func init(){
	fmt.Println("The init function has been triggered.")
}

func main(){
	//固定长度的数组,而且显示初始化
	arr :=[5]int{1,2,3,4,5}

	//固定长度数组，隐式初始化（默认是int的 false，也就是0，而不是c 语言那种随机内存)
	arr2  := [5] int {}

	printArray(&arr)

	printArray(&arr2)

	os.Exit(255)
}

func printArray(arr *[5]int){
	arrLen := len(*arr)

	for i:=0;i<arrLen;i++{
		fmt.Print(" ",(*arr)[i])
	}
}