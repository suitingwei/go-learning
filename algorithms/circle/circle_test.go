package circle

import (
	"fmt"
	"testing"
)

func lastRemaining(n int, m int) int {

	arr := make([]int, n)
	index := 0
	times := 1

	for {
		if times == n {
			break
		}

		step := 0

		for {
			if arr[index] == 0 {
				step++
			}

			if step == m {
				arr[index] = -1
				break
			}

			index++

			if index == n {
				index = 0
			}
		}
		fmt.Printf("%v\n", arr)
		times++
	}

	fmt.Printf("%v\n", arr)

	for i, val := range arr {
		if val == 0 {
			return i
		}
	}

	return 0
}

func TestLastRemaing(t *testing.T) {
	lastRemaining(5, 3)
}
