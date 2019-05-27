package panic

import (
	"fmt"
	"log"
)

func LearnPanic() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	dbProcess()

	fmt.Println("fuck you")
}

func dbProcess() {
	defer fmt.Println("close database handler")
	defer fmt.Println("close log ")

	panic("the fucking panic")
}
