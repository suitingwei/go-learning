package generics

import "testing"

type stringer interface {
	String() string
}

//字符类型累加器，把入参字符累加到内部的string中
type plusser interface {
	//讲入参累加到内部字符串中
	Plus(string)

	//获取最终累加结果
	Result() string
}

type Person struct{ Name string }

func (p Person) String() string { return p.Name }

//逗号加法器
type CommaPlusser struct {
	str string
}

func (p *CommaPlusser) Plus(s string) {
	if p.str == "" {
		p.str += s
	} else {
		p.str += "," + s
	}
}

func (p CommaPlusser) Result() string {
	return "normal plusser result=" + p.str
}

//可以声明一个interface作为type parameter
//本函数功能：将入参的args转换为string，然后累加到 result里，返回字符串
func stringPlus[S stringer, P plusser](plusser P, args []S) string {
	for i := 0; i < len(args); i++ {
		plusser.Plus(args[i].String())
	}
	return plusser.Result()
}

var john = Person{Name: "john"}
var alice = Person{Name: "alice"}
var bob = Person{Name: "bob"}

//取一个跟人没有任何关系的结构体，也代表着实际项目中，使用泛型的场景，都是完全不同的结构体
type File struct{ FileName string }

func (f File) String() string {
	return "__file__" + f.FileName
}

var zipFile = File{FileName: "movie.zip"}
var mp3File = File{FileName: "emmey.mp3"}
var picFile = File{FileName: "earth.png"}

//泛型的type parameter可以传递的不仅仅是type sets，还可以是interface
func TestInterface(t *testing.T) {
	normalPlusser := &CommaPlusser{}

	people := []Person{alice, bob, john}

	t.Logf("泛型[逗号加法器,人群]=%s\t", stringPlus[Person, plusser](normalPlusser, people))

	filePlusser := &CommaPlusser{}
	files := []File{zipFile, mp3File, picFile}
	t.Logf("泛型[逗号加法器,文件]=%s\t", stringPlus[File, plusser](filePlusser, files))

	//可以不传递泛型函数的构造初始化type parameters,编译器会进行类型推断
	t.Logf("泛型推断[逗号加法器,人群]=%s\t", stringPlus(normalPlusser, people))
}
