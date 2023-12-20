package leetcode_hot_100

import "strconv"

var wordMap = map[byte]bool{
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
}

// TODO again
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	i := 0
	for i < len(s) {
		if !wordMap[s[i]] {
			num, startIndex, endIndex := getNum(&s, i)
			res = res + multiStr(decodeString(s[startIndex:endIndex]), num)
			i = endIndex + 1
			continue
		}

		res = res + s[i:i+1]
		i++
	}

	return res
}

func multiStr(s string, num int) string {
	res := ""
	for i := 0; i < num; i++ {
		res = res + s
	}
	return res
}

func getNum(s *string, i int) (int, int, int) {
	var num int64
	var leftNum int
	var rightNum int
	var startIndex int
	var endIndex int
	for j := i; j < len(*s); j++ {
		if (*s)[j] == '[' {
			leftNum++
			if num == 0 {
				num, _ = strconv.ParseInt((*s)[i:j], 10, 64)
				startIndex = j + 1
			}
		}

		if (*s)[j] == ']' {
			rightNum++
			if leftNum == rightNum {
				endIndex = j
				break
			}
		}
	}
	return int(num), startIndex, endIndex
}
