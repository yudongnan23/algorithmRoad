package leetcode_hot_100

var bracketMap = map[byte]byte{
	41:  40,
	93:  91,
	125: 123,
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]byte, 0)

	remove := func(b byte) {
		v, ok := bracketMap[b]
		if ok && len(stack) > 0 && v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			return
		}

		stack = append(stack, b)
	}

	for i := 0; i < len(s); i++ {
		remove(s[i])
	}

	return len(stack) == 0
}
