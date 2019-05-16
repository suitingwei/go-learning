package _defer

import (
	"fmt"
	"time"
)

func LearnBasic() {
	//尽管defer会在函数结束的时候才会触发，但是他的传入参数早就做好了镜像了
	defer func(t time.Time) {
		fmt.Printf("The defer function triggered at :%s\n", t.Format("2006-01-02 15:04:05"))
	}(time.Now())

	//让这个函数在 5 秒后才结束
	time.Sleep(time.Second * 3)

	fmt.Printf("The function ended at:%s\n", time.Now().Format("2006-01-02 15:04:05"))

	output := []byte{1, 2, 3}

	newOutput := string(output[:])

	fmt.Println(newOutput)
}
