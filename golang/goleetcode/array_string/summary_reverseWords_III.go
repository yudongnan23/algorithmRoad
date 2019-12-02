/*
	给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序
*/

package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	sList := strings.Split(s, " ")
	sLength := len(sList)
	resultList := []string{}
	for index := 0; index < sLength; index++ {
		if sList[index] != "" {
			var curString string
			for _, val := range sList[index] {
				curString = string(val) + curString
			}
			resultList = append(resultList, curString)
		}
	}
	return strings.Join(resultList, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	ss := ReverseWords(s)
	fmt.Println(ss)
}
