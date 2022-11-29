package _2

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	result := []string{}
	dfs("", 0, 0, n, &result)
	return result
}

//leftCount 左括号的数量
//totalCount 总数量(对数)
func dfs(path string, leftCount int, rightCount int, totalCount int, result *[]string) {
	//递归重点，合法字符，加入结果集合
	if len(path) == totalCount*2 {
		*result = append(*result, path)
		return
	}

	//加入左括号,只要左括号还有，就可以用，不怕收不回去
	if leftCount < totalCount {
		dfs(path+"(", leftCount+1, rightCount, totalCount, result)
	}

	if leftCount > rightCount {
		//加入右括号,必须已经出现的左括号少，否则就关不上了
		dfs(path+")", leftCount, rightCount+1, totalCount, result)
	}
}

func TestName(t *testing.T) {
	result := generateParenthesis(3)
	for _, val := range result {
		fmt.Println(val)
	}
}
