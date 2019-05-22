package funcs

import (
	"fmt"
)

func main() {
	str := "hello world my name is sui"

	for index, value := range str {
		fmt.Printf("index=%d,value=%s\n", index, string(value))
	}

}
