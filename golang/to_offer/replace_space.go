package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace( s string ) string {
	if len(s) == 0 {
		return ""
	}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			s = s[0:i] + "%20" + s[i+1:]
		}
	}

	return s
}

func main () {
	fmt.Println(replaceSpace("We Are Happy"))
}