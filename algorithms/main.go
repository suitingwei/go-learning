package algorithms

import "fmt"

func main() {
	array := [][]int{
		{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1},
	}

	array = flipAndInvertImage(array)

	fmt.Println(array)
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
