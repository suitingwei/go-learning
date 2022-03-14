package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	name string
	age  int
}

type student struct {
	user
	school string
}

type myInt int

func TestReflectBasicUsage(t *testing.T) {
	var user1 = user{
		name: "xiaoming",
		age:  10,
	}
	var student1 = student{
		user:   user1,
		school: "MIT",
	}

	testcases := []interface{}{
		user1,
		1,
		myInt(1),
		int8(32),
		int16(22),
		32.11,
		"hello wolrd",
		[]int{1, 2, 3},
		student1,
	}

	for _, testCase := range testcases {
		printData(testCase)
		//kindVsName(testCase)
	}
}

//Kind 是类型,Name 是名字
//比如都是int,你可以起1w个别名，Name各不相同，但是Kind都是int
//或者说有很多struct，但是kind都是结构体
func kindVsName(i interface{}) {
	typeof1 := reflect.TypeOf(i)
	fmt.Printf("Kind=%s,Name=%s\n", typeof1.Kind(), typeof1.Name())
}

func printData(i interface{}) {

	typeof := reflect.TypeOf(i)
	valueof := reflect.ValueOf(i)

	switch typeof.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		{
			fmt.Println("It's a int type.")
		}
	case reflect.Float32, reflect.Float64:
		{
			fmt.Println("It's a float type.")
		}
	case reflect.Struct:
		{
			fmt.Println("It's a struct type.")
			for i := 0; i < typeof.NumField(); i++ {
				fmt.Println("\tField type=", valueof.Field(i).Kind())
				valueof.Field(i).Pointer()
				//structValue:= valueof.Field(i)
				//printData(structValue)
			}
		}
	case reflect.Slice:
		{
			elementType := typeof.Elem()
			fmt.Printf("It's a slice type, and the element type is=%s\n", elementType.Kind())
		}
	default:
		{
			fmt.Println("Unknown type.")
		}
	}
}
