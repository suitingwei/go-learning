package _2

import (
	"fmt"
	"testing"
)

func minDistance(word1 string, word2 string) int {

	dp := [][]int{}

	width := len(word1)
	length := len(word2)

	//初始化dp矩阵，多加了一个元素，对应dp[0][0]
	for i := 0; i <= width; i++ {
		dp = append(dp, make([]int, length+1))
	}

	for i := 1; i <= length; i++ {
		dp[0][i] = i
	}

	for j := 1; j <= width; j++ {
		dp[j][0] = j
	}

	for i := 1; i <= width; i++ {
		for j := 1; j <= length; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	fmt.Println(dp)
	return dp[width][length]
}

func min(nums ...int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func TestName(t *testing.T) {
	minDistance("act", "acg")
}
