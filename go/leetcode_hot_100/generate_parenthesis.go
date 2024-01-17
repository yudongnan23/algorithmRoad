package leetcode_hot_100

// TODO three
func generateParenthesis(n int) []string {
	res := make([]string, 0, 1000)
	var backtracking func(int, int, string)
	backtracking = func(left, right int, curStr string) {
		if right == 0 {
			res = append(res, curStr)
			return
		}

		if left > 0 {
			backtracking(left-1, right, curStr+"(")
		}

		if right > left {
			backtracking(left, right-1, curStr+")")
		}
	}

	backtracking(n, n, "")
	return res
}
