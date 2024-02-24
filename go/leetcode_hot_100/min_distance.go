package leetcode_hot_100

// TODO three
func minDistance(word1 string, word2 string) int {
	word1 = "-" + word1
	word2 = "-" + word2
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
		if i == 0 {
			for j := 0; j < m; j++ {
				dp[i][j] = j
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minInThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[n-1][m-1]
}

func minInThree(i, j, k int) int {
	if i < j && i < k {
		return i
	}

	if j < i && j < k {
		return j
	}

	return k
}
