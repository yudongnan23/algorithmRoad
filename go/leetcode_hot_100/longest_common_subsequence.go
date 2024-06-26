package leetcode_hot_100

// TODO three - 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	text1 = "-" + text1
	text2 = "-" + text2
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	longest := 0
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				}
			}

			if dp[i][j] > longest {
				longest = dp[i][j]
			}
		}
	}
	return longest
}
