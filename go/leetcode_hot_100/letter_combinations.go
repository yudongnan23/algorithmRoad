package leetcode_hot_100

import "math"

var keyMap = map[byte][]string{
	50: {"a", "b", "c"},
	51: {"d", "e", "f"},
	52: {"g", "h", "i"},
	53: {"j", "k", "l"},
	54: {"m", "n", "o"},
	55: {"p", "q", "r", "s"},
	56: {"t", "u", "v"},
	57: {"w", "x", "y", "z"},
}

// TODO again
func letterCombinations(digits string) []string {
	res := make([]string, 0, int(math.Pow(4, float64(len(digits)))))

	dfsIIIII(digits, "", &res)
	return res
}

func dfsIIIII(s string, curS string, res *[]string) {
	if len(s) == 0 {
		if len(curS) > 0 {
			*res = append(*res, curS)
		}
		return
	}
	for j := 0; j < len(keyMap[s[0]]); j++ {
		dfsIIIII(s[1:], curS+keyMap[s[0]][j], res)
	}
}
