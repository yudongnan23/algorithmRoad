/*

	给定两个二进制字符串，返回他们的和（用二进制表示）。
    输入为非空字符串且只包含数字 1 和 0。
 */
package main

import (
	"fmt"
)
import "math"
import "strconv"
import "strings"

func addBinary(a string, b string) string {
	if len(a) == 0{
		return b
	}
	if len(b) == 0{
		return a
	}
	if strings.EqualFold(a, "0") && strings.EqualFold(b, "0"){
		return "0"
	}
	var result string
	new_a := reverse(a)
	new_b := reverse(b)
	var int_a int64
	var int_b int64
	for i := 0; i < len(new_a); i++{
		if new_a[i:i+1] == "1"{
			int_a = int_a + int64(math.Pow(2, float64(i)))
		}
	}
	for j := 0; j < len(new_b); j++{
		if new_b[j:j+1] == "1"{
			int_b = int_b + int64(math.Pow(2, float64(j)))
		}
	}
	result = ten_to_two(int_a + int_b)
	return result
}
func ten_to_two(num int64) string {
	var result string
	for ;num >= 1; num/=2{
		result = strconv.Itoa(int(num)%2) + result
	}
	return result
}

func reverse(str string) string{
	var new_str string
	for i := 0; i < len(str); i ++{
		new_str = str[i:i+1] + new_str
	}
	return new_str
}

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))
}