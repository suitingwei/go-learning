package channel

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	queue := make(chan int)

	doubledChan := doubled(queue)

	go func() {
		for data := range doubledChan {
			fmt.Println(data)
		}
	}()

	for i := 0; i < 100; i++ {
		queue <- i
	}
}

func doubled(num <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for {
			result <- 2 * <-num
		}
	}()

	return result
}
