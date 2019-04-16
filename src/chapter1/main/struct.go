package main

import "fmt"

type user struct {
	name     string //ansidon
	age      int    //28
	location []int
}

type admin struct {
	id     int
	person user
}

/**
 * 这个方法相当于是 user 这个 struct 的 method.
 * 所以声明的时候，先指定了是哪一个结构体的 method，(u user)
 * 光这个语法就特么挺反人类的了，不像 c 在结构体里加函数指针。也不像java,php在 class 里声明 method
 * 这个func 关键字和 真正的函数名 toString之间的部分叫做函数接受者。
 * go语言里分为：值接受者，或者指针接受者。
 */
func (u user) toString() {
	fmt.Printf("Convert struct user into string:\n\tage=%d\n\tname=%s\n\tlocation=%v\n",
		u.age, u.name, u.location,
	)
}

/**
 * 修改结构体的数据（这里没有用指针来传参，所以需要看一下是否真的修改了结构体)
 */
func (u user) fakeChangeAge(newAge int) {
	u.age = newAge
}

/**
 * 使用指针传参，真的修改结构体
 * 只要声明了指针接受者，不论调用的时候是使用 user.realChangeAge(100);
 * 还是 (&user).realChangeAge(100);都能够调用这个方法。go 会自动转换非指针的到指针。
 */
func (u *user) realChangeAge(newAge int) {
	//作为指针调用，
	u.age = newAge

	//(*u).age = newAge
}

func main() {
	//basicStruct()

	structMethod()
}

func basicStruct() {

	list := []user{
		{name: "sui", age: 26, location: []int{1, 2, 3}},
	}

	adminUser := admin{
		id: 120,
		//嵌套结构体在初始化的时候，内部结构体必须也指定结构体的类型。
		//比如这个admin.user,初始化 user 的时候，不能直接用结构体初始化，而是必须指定下类型，再用值初始化
		//不能直接写 person: {"sui", 26, []int{1, 2, 3} }
		person: user{"sui", 26, []int{1, 2, 3}},
	}

	fmt.Printf("%v\n", adminUser.person)

	for index, userTemp := range list {
		fmt.Printf("index=%d,value=%v\n", index, userTemp)
	}

}

func structMethod() {

	user1 := user{name: "sui", age: 26, location: []int{1, 2, 3}}

	user1.toString()

	user1.fakeChangeAge(9899)

	user1.toString()

	user1.realChangeAge(9899)

	user1.toString()

}
