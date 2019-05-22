package goroutine

import (
	"fmt"
	"my-go-learnings"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter1 int64 //各个 routine 都要修改的资源
	wg1      sync.WaitGroup
)

/**
 * 资源竞争
 */
func LearnRareRoutinesWithLock() {

	my_go_learnings.wg.Add(2)

	fmt.Println("Create goroutines")
	go increaseCounter(1)
	go increaseCounter(2)

	fmt.Println("Waiting to finish")
	my_go_learnings.wg.Wait()
	fmt.Println("Final Counter", counter1)
}

func increaseCounter(id int) {
	defer my_go_learnings.wg.Done()

	for count := 0; count < 2; count++ {
		//原子的+1,相当于routine1修改 counter 的时候，routine2是不能读取更新的
		atomic.AddInt64(&counter1, 1)

		runtime.Gosched()
	}
}
