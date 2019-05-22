package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter2 int64 //各个 routine 都要修改的资源
	wg2      sync.WaitGroup
	mutex2   sync.Mutex
)

/**
 * 资源竞争
 */
func LearnRareRoutinesWithMutex() {
	wg2.Add(2)

	go increaseCounter2(1)
	go increaseCounter2(2)

	wg2.Wait()
	fmt.Println("Final Counter", counter2)
}

func increaseCounter2(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		//临界区，只有一个 goroutine 能进入这个区域，其他的只能阻塞
		mutex2.Lock()
		{
			value := counter2

			//即便当前线程退出去了，别的 goroutine 也不能去操作整个临界区的代码
			runtime.Gosched()

			value++

			counter2 = value
		}
		mutex2.Unlock()
	}
}
