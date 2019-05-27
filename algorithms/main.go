package algorithms

import (
	"fmt"
)

func main() {
	//array := [][]int{
	//	{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1},
	//}
	//
	//array = flipAndInvertImage(array)
	//
	//fmt.Println(array)

	result := largeGroupPositions("abcdddeeeeaabbbcd")

	fmt.Println(result)
}

func flipAndInvertImage(A [][]int) [][]int {
	var result [][]int
	for _, outer := range A {
		var tempArray []int
		for i := len(outer) - 1; i >= 0; i-- {
			if outer[i] == 1 {
				tempArray = append(tempArray, 0)
			} else {
				tempArray = append(tempArray, 1)
			}
		}
		result = append(result, tempArray)
	}

	return result
}

func largeGroupPositions(str string) [][]int {
	if len(str) < 3 {
		return [][]int{}
	}

	result := [][]int{}
	length := 1 // length is the longest group char length
	startIndex := 0

	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			//if the distance is greater or equals to 3, add it to the result slice
			if length >= 3 {
				result = append(result, []int{startIndex, i - 1})
			}

			//update the start and end index
			length, startIndex = 1, i
		} else {
			length++
		}
	}

	if length >= 3 {
		result = append(result, []int{startIndex, len(str) - 1})
	}

	return result
}
