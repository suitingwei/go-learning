package _00

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	islandCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j, &grid)
				islandCount++
			}
		}
	}
	return islandCount
}

//周围八个点
var directions = [][]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {0, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func dfs(i, j int, grid *[][]byte) {
	fmt.Printf("i=%d,j=%d,grid=%v", i, j, grid)
	(*grid)[i][j] = '0'
	for _, dir := range directions {
		newI := i + dir[0]
		newJ := j + dir[1]

		//圆心不管
		if newI == i && newJ == j {
			continue
		}

		if newI >= 0 && newI < len(*grid) && newJ >= 0 && newJ < len((*grid)[0]) {
			dfs(newI, newJ, grid)
		}
	}
}
