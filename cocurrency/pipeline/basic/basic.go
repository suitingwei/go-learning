package main

import (
	"fmt"
	"time"
)

func returnChanFunc(max int) <-chan int {
	result := make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Println("clock ticking", i)
			result <- i
		}
	}()

	return result
}

func main() {
	//for val := range returnChanFunc(100){
	//	fmt.Println("Main function got value",val)
	//}
	returnChanFunc(100)

	time.Sleep(time.Second * 5)
}
