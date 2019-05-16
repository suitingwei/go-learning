package reflection

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("The reflection's basic init function has been triggered!")
}

type X int

func LearnReflection() {
	a := X(100)

	t := reflect.TypeOf(&a)

	intVar := 10

	t2, t3 := reflect.TypeOf(intVar), reflect.TypeOf(&intVar)
	fmt.Println("Type of variable a is " + t.Name())

	fmt.Println(t2, t3)
	fmt.Println(t2.Kind(), t3.Kind())

	//用基础数据类型进行组装，这个时候注意，语法上是用基础类型的 typeof 结果作为参数
	//而不能直接用基础数据类型比如int
	t4 := reflect.ArrayOf(10, reflect.TypeOf(int(0)))

	//注意比较数组的时候，数组的长度也是参与比较的。而如果创建的变量是切片，那长度就是空的
	arrayInt := [8]int{1, 2, 3, 4, 5, 6}
	t5 := reflect.TypeOf(arrayInt)

	fmt.Println(t4, t5, t4 == t5)

	//5.ELem方法，判断array,slice,map,channel,ptr的元素的基础类型
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())           //map[string] => int
	fmt.Println(reflect.TypeOf([]int{}).Elem())                    //int[]
	fmt.Println(reflect.TypeOf(map[string][]int{}).Elem())         //map[string] => []int,嵌套了两次，map 里面是一个int slice。返回的就是内部的复合结构，[]int
	fmt.Println(reflect.TypeOf(map[string][]int{}).Elem().Elem())  //map[string] => []int,嵌套了两次，而 elem 也可以嵌套调用。
	fmt.Println(reflect.TypeOf(map[string][]*int{}).Elem().Elem()) //map[string] => []*int,嵌套了两次，而 elem 也可以嵌套调用。

	//6.Struct复杂结构
	type user struct {
		name string `field:"name" type:"varchar(50)"`
		age  int    `field:"age" type:"int"`
	}
	type manager struct {
		user  user
		title string
	}
	var m manager
	t6 := reflect.TypeOf(&m)
	fmt.Println("Type of manager is", t6, "kind of manager is", t6.Kind()) //typeof是manager*吧，kind应该是ptr

	//获取Elem，也就是基础类型，好奇这个是啥，毕竟里面是一个复杂的结构
	if t6.Kind() == reflect.Ptr {
		t6 = t6.Elem()
	}
	//日了狗了，就是 manager 这个结构体
	fmt.Println("Elem of manager is ", t6)

	for i := 0; i < t6.NumField(); i++ {
		f := t6.Field(i)
		fmt.Println("\tManager Field", f.Name, f.Type, f.Offset)

		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println("\t\tManager's ", f.Name, " field", af.Name, af.Type)
			}
		}
	}

	//7.struct tag

}
