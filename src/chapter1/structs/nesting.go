package structs

import "fmt"

type father struct {
	user   User
	belong string
}

type user struct {
	Name  string
	Email string
}
type PublicFather struct {
	user user
	Age  int
}

func (f father) earnMoney() {
	fmt.Println("The father is earning money...")
}

func LearnNestingStruct() {

	father := father{
		user:   User{"Jon", "jon@email.com"},
		belong: "earth mother",
	}

	sendNotification(father.user)

	father.user.notify()

	//作为嵌套结构体的外层结构体，可以直接调用内部的方法
	father.earnMoney()
}
