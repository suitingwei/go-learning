package sort

import "testing"

func TestSort(t *testing.T) {
	dummyData := []int{1, 14, 6, 2, 3, 7}
	//dummyData := []int{5, 1, 1, 2, 0, 0}

	t.Log(sortArray(dummyData))
}

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	left := 0
	right := len(nums) - 1

	if left >= right {
		return nums
	}

	middle := (right-left)/2 + left

	arr1 := mergeSort(nums[left : middle+1])
	arr2 := mergeSort(nums[middle+1 : right+1])

	return mergeSortedArr(arr1, arr2)
}

func mergeSortedArr(arr1, arr2 []int) []int {

	var result []int

	i := 0
	j := 0

	for {
		if i >= len(arr1) || j >= len(arr2) {
			break
		}
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	k := i
	remainingData := arr1

	if i == len(arr1) {
		k = j
		remainingData = arr2
	}

	for ; k < len(remainingData); k++ {
		result = append(result, remainingData[k])
	}

	return result
}
