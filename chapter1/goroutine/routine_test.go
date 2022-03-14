package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutineWithoutFuncParam(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 5000)
}
