/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

	有效字符串需满足：

		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。
 */


package main

import "fmt"

func isValid(s string)bool{
	valid := true

	openBrackets := []string{"(", "[", "{"}
	closeBrackets := []string{")", "]", "}"}
	stack := []string{}
	for i := 0; i < len(s); i++{
		curStr := s[i:i+1]
		if isInclude(openBrackets, curStr){
			stack = append(stack, curStr)
			continue
		}
		if  isInclude(closeBrackets, curStr){
			stackLength := len(stack)
			index := findIndex(closeBrackets, curStr)

			if stackLength == 0 || openBrackets[index] != stack[stackLength - 1]{
				valid = false
				break
			}

			stack = stack[0:stackLength - 1]
		}
	}

	if len(stack) != 0{
		valid = false
	}

	return valid
}

func findIndex(str []string, target string) int{
	var targetIndex int
	for index, value := range str{
		if value == target{
			targetIndex = index
		}
	}
	return targetIndex
}

func isInclude(str []string, target string) bool{
	include := false
	for _, value := range str{
		if value == target{
			include = true
		}
	}
	return include
}

func main() {
	s := "[{}]"
	fmt.Println(isValid(s))
}
