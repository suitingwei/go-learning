package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(generateOnce())
	}

}

func hashSlice(nums []int) string {
}
func generateOnce() []int {
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, rand.Intn(100))
	}
	return result
}
