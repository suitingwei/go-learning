package benchmark

import (
	"math/rand"
	"time"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func generateSliceWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	//提前设置cap内存大小
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generateSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	//没有设置cap内存大小
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
