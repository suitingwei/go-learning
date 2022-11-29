package _94

import (
	"strconv"
	"unicode"
)

func decodeString(s string) string {

	numStack := []string{}
	strStack := []string{}

	result := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if unicode.IsDigit(rune(char)) {
			if len(numStack) == 0 {
				numStack = append(numStack, string(char))
			} else {
				numStack[len(numStack)-1] += string(char)
			}
		} else if char == '[' {
			//如果是[符号，标识要进入新的一个level，那么两个栈的栈顶要移动

		} else if char == ']' {
			num := 1
			//准备生成答案，从数字和符号栈pop，如果数字是空，那么数字shi1
			if len(numStack) > 0 {
				num, _ = strconv.Atoi(numStack[len(numStack)-1])
			}

			str := strStack[len(strStack)-1]

			result += buildStr(str, num)

			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
		} else {
			//如果是字符串，那么一样拼接
			if len(strStack) == 0 {
				strStack = append(strStack, string(char))
			} else {
				strStack[len(strStack)-1] += string(char)
			}
		}
	}
	return result
}

func buildStr(str string, num int) string {
	result := ""

	for i := 0; i < num; i++ {
		result += str
	}
	return result
}
