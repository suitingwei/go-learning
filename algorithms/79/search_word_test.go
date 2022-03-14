package _9

import "testing"

var directions = [][2]int{
	{-1, 0}, //上
	{1, 0},  //下
	{0, -1}, //左
	{0, 1},  //右
}

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"

	t.Log(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	row := len(board)
	column := len(board[0])

	visited := map[[2]int]bool{}

	var check func(i, j, k int) bool
	//检查从当前地方搜索，是否有好结果
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		point := [2]int{i, j}
		visited[point] = true

		defer func() {
			visited[point] = false
		}()

		for _, dir := range directions {
			newRow := i + dir[0]
			newCol := j + dir[1]

			if newRow >= 0 && newRow < row && newCol >= 0 && newCol < column {
				newPoint := [2]int{newRow, newCol}
				if visited[newPoint] {
					continue
				}
				return check(newRow, newCol, k+1)
			}
		}
		return false
	}

	//从任何一个地方搜索
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
