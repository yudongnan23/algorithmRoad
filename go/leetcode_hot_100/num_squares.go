package leetcode_hot_100

var squareNums = make([]int, 0)
var dp = []int{0}

// TODO three
func numSquares(n int) int {
	if n < len(dp) {
		return dp[n]
	}
	for i := len(squareNums) + 1; i*i <= n; i++ {
		squareNums = append(squareNums, i*i)
	}
	for i := len(dp); i <= n; i++ {
		dp = append(dp, i)
		for _, square := range squareNums {
			if i < square {
				break
			}
			if dp[i-square]+1 < dp[i] {
				dp[i] = dp[i-square] + 1
			}
		}
	}
	return dp[n]
}
