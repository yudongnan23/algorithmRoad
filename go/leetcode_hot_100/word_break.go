package leetcode_hot_100

// TODO three - 单词拆分
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, str := range wordDict {
			if i-len(str) < 0 {
				continue
			}
			if dp[i-len(str)] && s[i-len(str):i] == str {
				dp[i] = dp[i-len(str)] && s[i-len(str):i] == str
				break
			}
		}
	}
	return dp[len(s)]
}
