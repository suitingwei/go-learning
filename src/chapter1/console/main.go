package main

import (
	"flag"
	"fmt"
)

func init() {
	fmt.Printf("init function in main package.")
}

func main() {

	var name string

	flag.StringVar(&name, "apiName", "", "Choose the api name you want")

	//定义了所有的 flags 之后，在调用Parse.这种带有状态的调用真是坑
	flag.Parse()

	fmt.Println("You have choose api:" + name)
}
