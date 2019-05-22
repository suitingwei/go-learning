package goroutine

import (
	"fmt"
	"my-go-learnings"
	"runtime"
	"sync"
)

func LearnLongRoutines() {

	var wg sync.WaitGroup

	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")

	wg.Wait()

	fmt.Println("Terminating program")
}

func printPrime(prefix string) {
	defer my_go_learnings.wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed!", prefix)
}
