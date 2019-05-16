package funcs

import (
	"fmt"
	"time"
)

func LearnTimer() {
	timer := time.NewTimer(time.Second * 3)

	<-timer.C

	fmt.Println("Event after certain seconds!")
}

func LearnTicker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
