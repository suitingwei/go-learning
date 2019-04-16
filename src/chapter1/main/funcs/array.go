package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{-4, -1, 0, 3, 10}

	result := []int{}
	for _, value := range arr {
		result = append(result, value*value)
	}
	fmt.Printf("%v\n", result)

	sort.Sort(sort.IntSlice(arr))

	fmt.Printf("%v\n", arr)

}
