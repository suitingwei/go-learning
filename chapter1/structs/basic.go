package structs

import "fmt"

//声明了一个接口
type notifier interface {
	notify()

	//如果 interface 里多了一个接口，而实现者没有全部实现，那么go 的编译器会认为这个实现者没有实现 interface，会报错
	//而不是将就着运行了
	//notify2()
}

type User struct {
	Name  string
	Email string
}

type admin struct {
	fullName   string
	familyName string
	email      string
}

//为啥这个方法叫notify就是接口了呢
//- 因为这个interface 就只有一个方法叫做 notify，所以实现了一个 interface 所有接口了，就是 implementation 了
//go 的结构体的是否public，完全看首字母是不是大写。是大写的就是公开，public， 否则就是 private
//因为这个interface 的方法是notify,所以这个大写的方法就不是接口了，而只是一个普通方法
func (u User) Notify() {
	fmt.Printf("Sending email to structUser:%s<%s>\n", u.Name, u.Email)
}

func (u User) notify() {
	fmt.Printf("\n\nSending 【!!!Secret!!!】email to structUser:%s<%s>\n", u.Name, u.Email)
}

func (admin admin) notify() {
	fmt.Printf("Sending email to structAdmin:%s-%s<%s>\n", admin.fullName, admin.familyName, admin.email)
}

func sendNotification(n notifier) {
	n.notify()
}
