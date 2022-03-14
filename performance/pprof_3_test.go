package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"testing"
	"time"
)

// main function
func TestRpf(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	//debug.SetGCPercent(-1)

	for i := 0; i < 10000; i++ {
		rpf()
	}

	time.Sleep(10 * time.Second)
}

// defining function having integer
// pointer as return type
func rpf() *string {

	// taking a local variable
	// inside the function
	// using short declaration
	// operator
	result := ""
	for i := 0; i <= 10000; i++ {
		result += strconv.Itoa(i)
	}

	// returning the address of lv
	return &result
}
