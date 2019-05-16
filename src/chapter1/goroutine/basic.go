package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func LearnGoRoutine() {
	//分配一个逻辑处理器给 goroutine 的调度器用
	runtime.GOMAXPROCS(1)

	//等待 goroutine 结束
	var wg sync.WaitGroup
	//add2表示等待两个goroutine结束
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//声明一个匿名函数来进行 goroutine
	go func() {

		//函数退出的时候调用wg.done()表示这个 routine 完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	/**
	 * 最终会先输出所有的大写字母，再输出小写字母。
	 * 因为这个大写字母的 routine 执行速度太快，所以还没来得及切换到另一个 routine，就完成了
	 */
	go func() {

		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting goroutines to finish!")

	wg.Wait()

	fmt.Println("\nTerminating program!")
}

func LearnRoutine() {
	fmt.Println("Start go routine")

	go func() {
		fmt.Println("Code in the go routine")
	}()

	//如果这里不 sleep，那么整个go 的 master 进程会结束，所有的协程都会被回收
	//time.Sleep(time.Second * 10)
	fmt.Println("Go routine ended")
}
