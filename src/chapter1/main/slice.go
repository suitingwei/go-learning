package main

import (
	"fmt"
)

func main(){

	//显示创建切片（动态数组)
	//slice := make([]int ,5)

	//partSlice();

	iterateSlice()

}

func appendSlice(){
	//nil切片
	slice3 :=[]int{}

	fmt.Println("Printing slice 3...")
	printSlice(&slice3)


	slice3 = append(slice3,50,50,20,10)

	fmt.Println("Printing slice 3 again...")
	printSlice(&slice3)
}

/**
 * slice[i:j:k],i表示起始切片的 index，j表示截止切片的 index（但是不包含j),k表示容量最后的index
 * slice[i:j:k] = nums[i,j)
 */
func partSlice(){

	//隐式创建切片，通过不指定数组长度
	slice2 :=[]int{1,2,3,4,5}

	//从切片中取出一个切片，这就是切片的名字来源,注意，这里的新切片没有复制一个额外的内存空间。
	//而是和原数组一样指向同一个数组，只不过 index和数量不一样
	//注意，slice 这个语法和 php 的 slice 是不一样的，他不会取j这个 index 新的元素
	partSlice2:=slice2[1:3]

	//printSlice(&partSlice2)
	//擦，go 居然本来就有直接打印 slice 的符号了，尴尬
	fmt.Printf("%v\n",partSlice2)

	//根据这个可以看出来，他只取出了index=2的元素
	partSlice2=slice2[2:3]
	printSlice(&partSlice2)

	//设置新切片的容量,k表示这个切片能达到的最后一个index
	partSlice2=slice2[3:4:4]
	printSlice(&partSlice2)
	fmt.Printf("old slice memory address:%p\n",&partSlice2)

	//如果添加元素的数量超过了切片的容量，那么会产生新的底层数组
	newPartSlice2 := append(partSlice2,888,999)

	fmt.Printf("new slice memory address:%p\n",&newPartSlice2)

	printSlice(&partSlice2)
}

func iterateSlice(){
	slice:=[]int{1,2,3,4,5,6,7}

	slice = append(slice,199,299)

	for index,value :=range slice{
		fmt.Printf("Index:%d\tvalue:%d\n",index,value)
	}
}

/**
 * 注意这里的s的类型，是一个不确定长度的int数组，在之前的代码中，如果传递的参数是定长数组，
 * 那么函数的参数列表声明的必须也是定长数组，只有这种slice 切片可以用不定长数组作为参数
 */
func printSlice(s *[]int){
	for i:=0;i<len(*s);i++{
		fmt.Print(" ",(*s)[i])
	}
	fmt.Println()
}
