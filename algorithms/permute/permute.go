package main

import "fmt"

func Permute(nums []int) [][]int {

	finalResult := [][]int{}

	gen(nums, map[int]bool{}, []int{}, &finalResult)

	return finalResult
}

func gen(nums []int, usedIndex map[int]bool, temp []int, finalResult *[][]int) {
	if len(temp) == len(nums) {
		fmt.Printf("Result Before= %v\n", *finalResult)
		newTemp := make([]int, len(temp))
		copy(newTemp, temp)
		*finalResult = append(*finalResult, newTemp)
		fmt.Printf("Result After= %v\n", *finalResult)
		return
	}

	for i := 0; i < len(nums); i++ {

		if _, ok := usedIndex[i]; ok {
			continue

		}

		temp = append(temp, nums[i])
		usedIndex[i] = true

		gen(nums, usedIndex, temp, finalResult)

		temp = temp[:(len(temp) - 1)]
		delete(usedIndex, i)
	}

	return
}

//[1,2,3,4]
//关键是回溯这一步
//[1]
// |____[1,2]
//        |_____[1,2,3]
//                  |_____[1,2,3,4]
//        |_____[1,2,4]
//                  |_____[1,2,4,3]
// |____[1,3]
// |____[1,4]
