package generics

import (
	"testing"
)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

//简单的数字比较
//泛型比较大小,这里有一个坑，comparable这个Constraint 其实应该叫做 equivalents 因为他只包含 ==, != 的判断
//不知道大小比较，导致在这里浪费很久,see https://github.com/golang/go/issues/51861,所以这里应该替换为
//func minGenerics[T comparable](a, b T) T {
func minGenerics[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

//也支持不定长参数
func minInSlice[T Ordered](args ...T) T {
	//这种类型声明怎么处理? :=的快捷语法是不支持的
	//result := T{} //failed:syntax error

	//相反，用普通的var 语法是支持的，目前看起来真是因为编译器不支持原因
	//那么这个type T的最小值是多少呢? 是每一个类型的默认值吧，也就是0?
	var result T
	if len(args) == 0 {
		return result
	}

	//总是默认空值，因为任何数字类型的default都是0，如果是
	//fmt.Printf("the default value of type T is =%v\n", result)

	//所以算法上来说，要用数组第0个元素赋初始值
	result = args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < result {
			result = args[i]
		}
	}
	return result
}

//类型别名也是支持的,ProductId虽然是int32别名，但是也继承自 Signed， 所以也是 Ordered
type ProductId int32

func TestMin(t *testing.T) {
	t.Logf("None Generic Min Func return =%v", min(10.11, 22.332))
	//如果是已知的类型，那么不需要声明 type parameter
	t.Logf("Generic Min Func return =%v", minGenerics(ProductId(11), ProductId(32)))

	//也可以显示声明type parameter
	t.Logf("Generic Min Func return =%v", minGenerics[ProductId](ProductId(11), ProductId(32)))
}

func TestMinInSlices(t *testing.T) {
	if result := minInSlice(1, 2, 3, 4, 5); result != 1 {
		t.Fatalf("want %v, got %v", 1, result)
	}
	if result := minInSlice(1, 2, -3, 4, 5); result != -3 {
		t.Fatalf("want %v, got %v", 1, result)
	}
}
