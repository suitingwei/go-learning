package main

import (
	"fmt"
)

func main() {

	//显示创建切片（动态数组)
	//slice := make([]int ,5)

	//partSlice();

	//iterateSlice()

	//multipleSlice()

	mapLearnings()
}

func appendSlice() {
	//nil切片
	slice3 := []int{}

	fmt.Println("Printing slice 3...")
	printSlice(slice3)

	slice3 = append(slice3, 50, 50, 20, 10)

	fmt.Println("Printing slice 3 again...")
	printSlice(slice3)
}

/**
 * slice[i:j:k],i表示起始切片的 index，j表示截止切片的 index（但是不包含j),k表示容量最后的index
 * slice[i:j:k] = nums[i,j)
 */
func partSlice() {

	//隐式创建切片，通过不指定数组长度
	slice2 := []int{1, 2, 3, 4, 5}

	//从切片中取出一个切片，这就是切片的名字来源,注意，这里的新切片没有复制一个额外的内存空间。
	//而是和原数组一样指向同一个数组，只不过 index和数量不一样
	//注意，slice 这个语法和 php 的 slice 是不一样的，他不会取j这个 index 新的元素
	partSlice2 := slice2[1:3]

	//printSlice(&partSlice2)
	//擦，go 居然本来就有直接打印 slice 的符号了，尴尬
	fmt.Printf("%v\n", partSlice2)

	//根据这个可以看出来，他只取出了index=2的元素
	partSlice2 = slice2[2:3]
	printSlice(partSlice2)

	//设置新切片的容量,k表示这个切片能达到的最后一个index
	partSlice2 = slice2[3:4:4]
	printSlice(partSlice2)
	fmt.Printf("old slice memory address:%p\n", &partSlice2)

	//如果添加元素的数量超过了切片的容量，那么会产生新的底层数组
	newPartSlice2 := append(partSlice2, 888, 999)

	fmt.Printf("new slice memory address:%p\n", &newPartSlice2)

	printSlice(partSlice2)
}

func iterateSlice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("The capacity of the slice is : ", cap(slice))

	slice = append(slice, 199, 299)

	for index, value := range slice {
		fmt.Printf("Index:%d\tvalue:%d\n", index, value)
	}

	//不需要index,只需要value。这语法其实有点类似php, list(,$value) = [1,2];
	for _, value := range slice {
		fmt.Printf("None index, value:%d\n", value)
	}

	//slice扩容是成倍扩展的，第一次是7，后来扩容是14
	fmt.Println("The capacity of the slice is : ", cap(slice))

}

/**
 * 测试多维数组
 */
func multipleSlice() {
	//多维数组
	arr := [][]int{{1002, 3912, 301092}, {2, 3, 4}}

	for _, firstDimension := range arr {
		for _, secondDimension := range firstDimension {
			fmt.Printf(" %d", secondDimension)
		}
		fmt.Println()
	}
}

/**
 * 注意这里的s的类型，是一个不确定长度的int数组，在之前的代码中，如果传递的参数是定长数组，
 * 那么函数的参数列表声明的必须也是定长数组，只有这种slice 切片可以用不定长数组作为参数
 */
func printSlice(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Print(" ", s[i])
	}
	fmt.Println()
}

func mapLearnings() {
	//普通的颜色映射 key => value
	colors := map[string]string{}

	//声明之后赋值。也是 map 中新增元素的方式
	colors["red"] = "apple"
	colors["yellow"] = "apple"
	colors["blue"] = "strawberry"

	colors2 := map[string]string{
		"Red":    "Apple",
		"Yellow": "Banana",
	}

	fmt.Printf("%v", colors2)

	for key, value := range colors {
		fmt.Printf("Key:%s\tValue:%s\n", key, value)
	}
	fmt.Println()

	colors3 := map[string][]string{"Red": {"Apple", "Strawberry"}}
	fmt.Printf("%v\n", colors3)

	//从 map 取出来的时候，可以多一个变量来判断这个 key 是否存在
	value, exists := colors3["Red"]

	if exists {
		fmt.Printf("Red key exists in the color3 map:%v, and value is:%v\n", colors3, value)
	}

	//如果 key 不在 map 中存在，那么返回值是一个空值,整个 map 的 value 是啥类型，返回的就是啥类型
	value, exists = colors3["Black"]

	if !exists {
		fmt.Printf("Key:%s not exists in color3 map:%v, and return value is:%v\n", "Black", colors3, value)
	}

	//删除 map 的一个 key
	delete(colors3, "Red")

	fmt.Printf("After deletetion, the color3 map is:%v\n", colors3)

	//直接复制一个字符串切片
	colors3["Black"] = []string{"new", "ballo"}

	fmt.Printf("After addtion, the color3 map is:%v\n", colors3)
}
