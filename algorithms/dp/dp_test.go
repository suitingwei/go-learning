package dp

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/perfect-squares/
// 完全平方数
func Test_perfect_squares(t *testing.T) {
	cases := map[int]int{
		4:   1,
		12:  3,
		50:  2,
		120: 3,
	}

	for key, value := range cases {

		if value != numSquares(key) {
			t.Errorf("N:%d should has result:%d, but got: %d\n", key, value, numSquares(key))
		}
	}

}

func numSquares(n int) int {
	candidates := findSquares(n)

	fmt.Printf("Find candicates:%v\n", candidates)

	//result  := []int{}
	//
	//findByBackTrace(candidates,n,&result,[]int{},make(map[string]bool))

	//fmt.Printf("Final result with length:%d \t%v\n",len(result),result)
	//return len(result)

	return findByDP(candidates, n)
}

func findByDP(candidates []int, target int) int {
	newCandidates := make([]bool, target+1)

	for _, value := range candidates {
		if value == target {
			return 1
		}
		newCandidates[value] = true
	}

	//1是完全平方数，最多用N个1来拼接,也就是N个步骤
	max := target

	dp := make([][]bool, max)

	//第一行的可达到的和，就是当前的这些备选值
	dp[0] = newCandidates

	for i := 0; i < max; i++ {
		// 当前的和有这么多情况：0,1,2,3,4,5...n,总共是n+1
		dp[i] = newCandidates
	}

	//printDP(dp)

	//i表示用的步数
	for i := 1; i < max; i++ {
		for j := target; j >= 0; j-- {
			//如果之前数字点亮了
			if dp[i-1][j] {
				for _, candidate := range candidates {
					if j+candidate <= target {
						dp[i][j+candidate] = true

						if j+candidate == target {
							return i + 1
						}
					}
				}
			}
		}
	}
	//printDP(dp)
	return max
}

func printDP(bools [][]bool) {
	for i, outer := range bools {
		fmt.Printf("[%d]:\t", i)
		for j, inner := range outer {
			fmt.Printf("[%d]:%t ", j, inner)
		}
		fmt.Println()
		fmt.Println()
	}
}

//使用回溯找到所有的结果
// candidates 备选的数字，从小到大排序
// target 目标数字
// result 所有可能的组成
// temp  当前的尝试路径
func findByBackTrace(candidates []int, target int, result *[]int, temp []int, memory map[string]bool) {
	currentSum := sumArray(temp)
	//找到了目标
	if currentSum == target {
		if len(*result) == 0 || len(temp) < len(*result) {
			*result = temp
			return
		}
	}

	key := arrayToKey(temp)

	if _, ok := memory[key]; ok {
		return
	}

	//尝试每一个数字
	for _, value := range candidates {
		//如果这个数字可以被加入路径
		if currentSum+value <= target {
			memory[key] = true
			findByBackTrace(candidates, target, result, append(temp, value), memory)
		}
	}
}

func arrayToKey(ints []int) string {
	sort.Ints(ints)

	return strings.Join(strings.Split(fmt.Sprint(ints), " "), "-")
	//return strconv.Itoa(sumArray(ints))
}

func sumArray(nums []int) int {
	result := 0
	for _, value := range nums {
		result += value
	}
	return result
}

//找到<=n的所有完全平方数
func findSquares(n int) []int {
	half := math.Sqrt(float64(n))

	result := []int{}
	for i := 1.0; i <= half; i++ {
		result = append(result, int(i*i))
	}
	return result
}
