package log

import (
	"fmt"
	"log"
	"os"
)

func init() {
	//通过打开文件来指定log 文件
	file, err := os.OpenFile("./log.file", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		fmt.Println("Error occurred to create log file!")
		os.Exit(1)
	}
	//设置输出 output
	log.SetOutput(file)

	//设置前缀
	log.SetPrefix("LEARN_LOG ")

	//使用位运算来设置 log 的各个标志位
	//使用位运算这个太特么聪明了，某一个标志是0000001,然后其他的是依次*2，对于2进制来说就是 0000010之类的
	//这样最终的结果就是位运算的或操作符。
	log.SetFlags(log.Ldate | log.Llongfile | log.Ltime)
}
func LearnLog() {

	//Log包的函数是支持多进程写的，他会通过临界区加锁
	log.Println("message 1")

	//这个方法会调os.Exit(1)
	log.Fatalln("Fatal message 1!")

	//这个方法会调用panic()
	//log.Panicln("Panic message 1!")
}
