package main

import (
	"fmt"
	"time"
)

func returnChanFunc(max int) <-chan int {
	result := make(chan int)

	//这个设计真的很灵性，首先result channel是逃逸了堆栈，从函数返回了，但是内存还是存在的
	//其次，开了一个额外的goroutine,来进行数据的写入，但是这个地方用了一个阻塞channel
	//所以必须等待别的goroutine读取数据，最终每次读取数据，才会在这里写入channel
	go func() {
		for i := 0; i < max; i++ {
			fmt.Println("clock ticking", i)
			//阻塞channel必须有人读取数据，才会进行下一次的写入操作
			result <- i
		}
		//这里最终必须关闭channel，否则会出现channel死锁。因为没有人写入channel了，但是main goroutine还在等待读取
		close(result)
	}()

	return result
}

func main() {
	for val := range returnChanFunc(100) {
		fmt.Println("Main function got value", val)
	}

	time.Sleep(1 * time.Second)
}
