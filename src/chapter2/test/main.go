package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess can you get a reward....")

	time.Sleep(time.Millisecond * 500)

	rand.Seed(int64(time.Now().Nanosecond()))
	random := rand.Intn(100)

	if random > 30 {
		fmt.Printf("You have won the reward:%d!\n", random)
	} else {
		fmt.Println("Sorry buddy, you will get it next time!")
	}
}
