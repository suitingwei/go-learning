package main

import "fmt"

type position struct {
	row int
	col int
}

//方向转换
//map key 的第一个值是当前方向，第二个值是要执行的转向指令
//map value是这次执行完之后的方向
var directions = map[[2]uint8]uint8{
	//北风
	[2]uint8{'N', 'L'}: 'W',
	[2]uint8{'N', 'R'}: 'E',

	//南风
	[2]uint8{'S', 'L'}: 'E',
	[2]uint8{'S', 'R'}: 'W',

	//西风
	[2]uint8{'W', 'L'}: 'S',
	[2]uint8{'W', 'R'}: 'N',

	//东风
	[2]uint8{'E', 'L'}: 'N',
	[2]uint8{'E', 'R'}: 'S',
}

func main() {

	path := "LGLGLGLGLGLGLG"

	arr := [100][100]bool{}

	arr[0][0] = true

	printRoute(arr)

	var direction uint8 = 'N'
	pos := position{0, 0}

	for i := 0; i < len(path); i++ {
		action := path[i]
		//如果不是前进指令，那么就是方向指令，这个时候根据专项指令集
		//计算出来新的方向
		if action != 'G' {
			direction = directions[[2]uint8{direction, action}]
		} else {
			switch direction {
			case 'N':
				arr[pos.row][pos.col+1] = true
			case 'S':
				arr[pos.row][pos.col-1] = true
			case 'W':
				arr[pos.row-1][pos.col] = true
			case 'E':
				arr[pos.row+1][pos.col] = true
			}
		}

		printRoute(arr)
	}
}

func printRoute(route [100][100]bool) {
	for _, row := range route {
		for _, col := range row {
			if col {
				fmt.Printf("*")
			} else {
				fmt.Printf("_")
			}
		}

		fmt.Println()
	}
}
