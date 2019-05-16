package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var ballWg sync.WaitGroup

func LearnRoutineWithChannel() {

	//整数类型，无缓冲通道
	//unbuffered:=make(chan int)

	//字符串类型，10字节缓冲的通道
	court := make(chan int)

	ballWg.Add(2)

	go player("Sui", court)

	go player("Wu", court)

	//发球
	court <- 1

	ballWg.Wait()
}

func player(name string, court chan int) {

	defer ballWg.Done()

	for {
		//从 channel 里读取数据的操作是阻塞的，直到通道里有数据之后才会返回
		//并且这个操作会锁住整个goroutine(啥意思，线程从逻辑调度器隔离？）
		ball, ok := <-court

		//如果通道已经关闭，那么ok返回 false
		if !ok {
			fmt.Printf("Player :%s won!\n", name)
			return
		}

		//随机数表示是否丢球
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s missed!\n", name)

			//关闭通道，表示我们输了
			close(court)

			return
		}

		//显示击球数，并且把数目加1
		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++

		//把球传过去,这个操作也是阻塞的，会锁住两个交换的 goroutine，直到交换完成。
		//如果是三个 goroutine 呢，这个数量是怎么定义的？
		court <- ball
	}
}
