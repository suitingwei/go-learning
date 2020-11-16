package _406

import (
	"fmt"
	"sort"
)

func rebuild(people [][]int) [][]int {

	//sort people by height
	sort.Slice(people, func(i, j int) bool {
		//if the height are same, compare by the k, the smaller k will be put forwards
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		//compare people by the height
		return people[i][0] < people[j][0]
	})

	fmt.Printf("people after sorted is %v\n", people)

	//the result container
	result := make([][]int, len(people))

	fmt.Printf("the result container is:%v\n", result)

	//rebuild the queue
	for i := 0; i < len(people); i++ {
		//the person count whose height is less than k
		count := 0
		//put the value from the beginning to the end
		for j := 0; j < len(result); j++ {
			// if this position is the index to place the person
			if result[j] == nil && count == people[i][1] {
				result[j] = people[i]
				fmt.Printf("We have found the position to place person[%d,%d] at index %d, now the result is%v\n",
					people[i][0], people[i][1], count, result)
				break
			}

			//if this place is empty or this person's height is less or equals than the current person, is should be count
			if result[j] == nil || result[j][0] >= people[i][0] {
				fmt.Printf("Result[%d] is %v\n", j, result[j])
				count++
			}
		}
	}
	return result
}
