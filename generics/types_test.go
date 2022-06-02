package generics

import (
	"fmt"
	"testing"
)

//测试泛型类型，是的，泛型不仅仅支持泛型函数，还支持泛型type

//泛型vector，其实是一个数组，但是其数组元素是一个泛型T
type Vector[T any] []T

//Generic types can have methods.
//泛型类型类似Vector也可以有自己的方法
//The receiver type of a method must declare the same number of type parameters as are declared in the receiver type's definition.
//方法接受者必须定义同样数量的type parameter(也就是 Vector[T any] 中的 T
//They are declared without any constraint.
//不过这些type parameter他们不需要再指定类型,这个时候名字对不上都无所谓
//泛型Vector的长度，没什么好说的，直接len即可,因为builtin function len本身支持泛型
func (v Vector[T]) Len() int  { return len(v) }
func (v Vector[S]) Len2() int { return len(v) }

//添加元素到Vector，这里要注意pointer receiver function，其实就是调用该函数的时候，传入了当前实例的指针
func (v *Vector[T]) Push(x ...T) {
	*v = append(*v, x...)
}

//POP will pop the first element of the vector
func (v *Vector[T]) Pop() T {
	if v.Len() == 0 {
		var result T
		return result
	}

	//get the first element of the vector
	result := (*v)[0]

	//pop the first element of the vector
	temp := (*v)[1:]

	//reassign the vector
	*v = temp
	return result
}

func (v Vector[T]) Print() {
	fmt.Printf("\nCurrent length of vector: %d\n", v.Len())
	for i := 0; i < len(v); i++ {
		fmt.Printf("%v,", v[i])
	}
	fmt.Println()
}

func TestVector(t *testing.T) {
	intVector := Vector[int]{1, 2, 3}

	intVector.Push(100)
	intVector.Push(3920)
	intVector.Push(29)

	intVector.Print()

	//增加泛型测试
	personVector := Vector[Person]{john, alice, bob}
	personVector.Print()
	personVector.Push(Person{Name: "tom"}, Person{Name: "fried"})
	personVector.Print()

	//测试pop
	topPerson := personVector.Pop()

	fmt.Printf("first element in the vector is=%v\n", topPerson)

	personVector.Print()
}

//类似泛型函数声明，先声明类型的名字,再声明泛型参数，最终声明泛型的结构(数组，结构体)
//func generics[T any] (arg T)

type Container[S stringer, P plusser] struct {
	stringHandler S
	plusser       map[string]P     //这里可以类似原来的形式使用泛型类型
	next          *Container[S, P] //支持递归(自我引用),可以生成类似链表的结构
}

//当创建函数的时候，不需要再添加type parameter的具体类型，比如以下声明是错误的
//func (c Container[S stringer,P plusser])  FuncName() string{}
//而是直接使用type parameter即可
func (c Container[S, P]) GetFinalName() string {
	c.stringHandler.String()
	return ""
}

type Iterator[K, V any] struct {
}

//func (c Iterator[K, V]) Next() (K, V, bool) {
//	return nil, nil, false
//}

//函数无法再声明type parameter了
