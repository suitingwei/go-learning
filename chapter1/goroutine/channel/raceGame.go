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

var raceGameWg sync.WaitGroup

func LearnChannelWithRaceGame() {
	//创建无缓冲的通道
	baton := make(chan int)

	//为最后一个跑步者吧计数器加一
	raceGameWg.Add(1)

	go Runner(baton)

	baton <- 1

	raceGameWg.Wait()
}

//runner 方法模拟比赛中的一个跑步者
func Runner(baton chan int) {

	var newRunner int

	//等待接力棒
	runner := <-baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1

		fmt.Printf("Runner %d to the line\n", newRunner)

		go Runner(baton)
	}

	//围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	//比赛是否结束
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over\n", runner)

		raceGameWg.Done()

		return
	}

	fmt.Printf("Runner %d Exchange with Runner %d\n", runner, newRunner)

	baton <- newRunner
}
