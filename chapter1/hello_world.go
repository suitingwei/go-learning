package main

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("The init function has been triggered.")
}

func main() {

	pool := sync.Pool{
		New: func() interface{} {
			return 1
		},
	}

	fmt.Println(pool.Get()) //1

	pool.Put(100)
	fmt.Println(pool.Get()) //100

	pool.Put(200)
	fmt.Println(pool.Get()) //200

	pool.Put(300)
	fmt.Println(pool.Get()) //200
}
