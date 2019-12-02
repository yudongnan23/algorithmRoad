/*
	给定一个字符串，逐个翻转字符串中的每个单词
*/

package main

import (
	"fmt"
	"strings"
)

// O(logn)
//func reverseWords(s string) string{
//	var resultString string
//	if s == ""{
//		return resultString
//	}
//	new_s := strings.Trim(s, " ")
//	wordsList := strings.Split(new_s, " ")
//	for index := len(wordsList) - 1; index >= 0; index--{
//		if resultString == ""{
//			resultString = resultString + wordsList[index]
//		}else if wordsList[index] != ""{
//			resultString = resultString + " " + wordsList[index]
//		}
//	}
//	return resultString
//}

// O(n)
//func reverseWords(s string) string{
//	var resultString string
//	if s == ""{
//		return resultString
//	}
//	for index := len(s) - 1; index >= 0;{
//		if string(s[index]) != " "{
//			var curString string
//			for ; index >= 0; index--{
//				if string(s[index]) != " "{
//					curString = string(s[index]) + curString
//				}else{
//					break
//				}
//			}
//			if resultString == ""{
//				resultString =  resultString + curString
//			}else{
//				resultString = resultString + " " + curString
//			}
//			index--
//		}else{
//			index--
//		}
//	}
//	return resultString
//}

func reverseWords(s string) string {
	wordsList := strings.Split(s, " ")
	resultString := []string{}
	for index := len(wordsList) - 1; index >= 0; index-- {
		if wordsList[index] != "" {
			resultString = append(resultString, wordsList[index])
		}
	}
	return strings.Join(resultString, " ")
}

func main() {
	s := "hello world"
	fmt.Println(reverseWords(s))
}
