package leetcode_hot_100

// TODO three - 最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	ans := ""
	for r := 0; r < n; r++ {
		for l := 0; l <= r; l++ {
			if r-l == 1 {
				dp[l][r] = s[l] == s[r]
			} else if r > l {
				dp[l][r] = (s[l] == s[r]) && dp[l+1][r-1]
			}
			if dp[l][r] && r-l+1 > len(ans) {
				ans = s[l : r+1]
			}
		}
	}
	return ans
}
