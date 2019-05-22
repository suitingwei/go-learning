package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int //各个 routine 都要修改的资源
	wg      sync.WaitGroup
)

/**
 * 资源竞争
 */
func LearnRareRoutines() {

	wg.Add(2)

	fmt.Println("Create goroutines")
	go increCounter(1)
	go increCounter(2)

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Final Counter", counter)
}

func increCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		//当前goroutine从线程退出
		runtime.Gosched()

		//修改本地 value
		value++

		//重新赋值
		counter = value
	}
}
