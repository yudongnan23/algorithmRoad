/*
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""
 */


package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string{
	if len(strs) == 0{
		return ""
	}
	minLength := min_lenth(strs)
	for i := 0 ; i < minLength; i++{
		for j := 0; j < len(strs); j ++{
			if strings.EqualFold(strs[j][i:i+1], strs[0][i:i+1]) == false{
				return strs[j][0:i]
			}
		}
	}
	return strs[0][0:minLength]

}

func min_lenth(strs []string) int{
	if len(strs) == 0{
		return 0
	}
	var result = len(strs[0])
	for _, str := range strs{
		if len(str) < result{
			result = len(str)
		}
	}
	return result
}

func main() {
	strs := []string{"flower", "flow", "flowight"}
	fmt.Println(longestCommonPrefix(strs))
}