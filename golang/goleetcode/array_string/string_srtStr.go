/*

	实现 strStr() 函数。
	给定一个 haystack 字符串和一个 needle 字符串，
	在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
	如果不存在，则返回  -1。
 */
package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0{
		return 0
	}
	len_needle := len(needle)
	for i := 0; i < len(haystack); i++{
		if len(haystack) - i < len_needle{
			break
		}else if strings.EqualFold(haystack[i:i+len_needle], needle){
			return i
		}
	}
	return -1

}

func main() {
	ss := "aaaaa"
	s := "bba"
	fmt.Println(strStr(ss, s))
}

