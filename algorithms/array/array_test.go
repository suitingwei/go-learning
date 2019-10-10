package array

import (
	"fmt"
)

func numIslands(matrix [][]byte) int {

	exploredPoints := make(map[[2]int]bool)

	maxWidth := len(matrix)

	if maxWidth == 0 {
		return 0
	}

	maxLength := len(matrix[0])

	result := 0

	for i := 0; i < maxWidth; i++ {
		for j := 0; j < maxLength; j++ {
			//如果这个点不是1
			if matrix[i][j] != '1' {
				continue
			}

			point := [2]int{i, j}

			//如果这个点已经探索过
			if _, ok := exploredPoints[point]; ok {
				continue
			}

			//找到以这个点为坐标点的岛屿
			islandMap := make(map[[2]int]bool)
			islandMap[point] = true
			findOuterLand(matrix, point, exploredPoints, maxWidth, maxLength)

			result++
		}
	}

	return result
}

//找到某一个点周围的所有土地
//matrix [][]int,地图
//point [][]int 当前探索的点。用一个长度为2的数组保存横纵坐标
//map[string]bool,保存已经探索过的点，这些探索过的点，在整个探索过程中不会再用到
func findOuterLand(matrix [][]byte, point [2]int, exploredPoints map[[2]int]bool, maxX, maxY int) {

	x, y := point[0], point[1]

	//以下几个点，都是针对人类眼中的上下左右
	leftPoint := [2]int{x, y - 1}
	rightPoint := [2]int{x, y + 1}
	upPoint := [2]int{x - 1, y}
	downPoint := [2]int{x + 1, y}

	var (
		hasUp    = true
		hasDown  = true
		hasLeft  = true
		hasRight = true
	)

	if x <= 0 {
		hasUp = false
	}
	if x >= maxX-1 {
		hasDown = false
	}

	if y <= 0 {
		hasLeft = false
	}

	if y >= maxY-1 {
		hasRight = false
	}

	if hasUp && matrix[upPoint[0]][upPoint[1]] == '1' {
		if _, outOk := exploredPoints[upPoint]; !outOk {
			exploredPoints[upPoint] = true
			findOuterLand(matrix, upPoint, exploredPoints, maxX, maxY)
		}
	}

	if hasDown && matrix[downPoint[0]][downPoint[1]] == '1' {
		if _, ok := exploredPoints[downPoint]; !ok {
			exploredPoints[downPoint] = true
			findOuterLand(matrix, downPoint, exploredPoints, maxX, maxY)
		}
	}

	if hasLeft && matrix[leftPoint[0]][leftPoint[1]] == '1' {
		if _, ok := exploredPoints[leftPoint]; !ok {
			exploredPoints[leftPoint] = true
			findOuterLand(matrix, leftPoint, exploredPoints, maxX, maxY)
		}
	}

	if hasRight && matrix[rightPoint[0]][rightPoint[1]] == '1' {
		if _, ok := exploredPoints[rightPoint]; !ok {
			exploredPoints[rightPoint] = true
			findOuterLand(matrix, rightPoint, exploredPoints, maxX, maxY)
		}
	}

	return
}
func printIsland(matrix [][]string, islandMap map[[2]int]bool) {

	maxY := len(matrix)
	maxX := len(matrix[0])

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if _, ok := islandMap[[2]int{i, j}]; ok {
				fmt.Print(" * ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
