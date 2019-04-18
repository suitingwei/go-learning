package funcs

import (
	"fmt"
	"sort"
)

func learnArray() {
	arr := []int{1, 3, 5, 6}

	result := searchInsert(arr, 5)

	fmt.Println(result)
}
func sortArray() {

	arr := []int{-4, -1, 0, 3, 10, 2}

	var result []int

	for _, value := range arr {
		result = append(result, value*value)
	}
	fmt.Printf("%v\n", result)

	//直接排序数组
	sort.Ints(arr)

	//sort.Sort(sort.IntSlice(arr))

	fmt.Printf("%v\n", arr)

}

func searchInsert(nums []int, target int) int {
	fmt.Println(len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
