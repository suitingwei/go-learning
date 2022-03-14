package main

import (
	"fmt"
	"testing"
	"time"
)

func TestUpdateDriverRpc(t *testing.T) {
	updateDriver("650911060392918", "0", 1)
}

func TestUpdateDriverFromCsv(t *testing.T) {
	//clean all driver attributes
	updateCity("./test.csv", hallQuickSwitchOff)
	//sleep for a while,in case triggering rate limiter of the uranus system
	time.Sleep(time.Second)

	//update all driver list
	updateCity("./test.csv", hallQuickSwitchOn)
}

func TestName(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		fmt.Printf("currentProcessingLine=%d\r", i)
	}
}
