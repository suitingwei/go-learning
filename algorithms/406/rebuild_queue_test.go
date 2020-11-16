package _406

import "testing"

func TestRebuild(t *testing.T) {

	people := [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}

	result := rebuild(people)

	t.Logf("The input value is %v\n", people)
	t.Logf("The result is %v\n", result)
}
