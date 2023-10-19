package to_offer

var _len = map[int]int{
	0: 1,
	1: 1,
	2: 1,
	3: 2,
	4: 4,
	5: 6,
	6: 9,
}

func cuttingBamboo(length int) int {
	if length <= 6 {
		return _len[length]
	}

	dp := make([]int, length+1)
	for i := 0; i <= 6; i++ {
		dp[i] = _len[i]
	}

	for i := 7; i <= length; i++ {
		dp[i] = maxLength(dp[i-2]*2, dp[i-3]*3)
	}

	return dp[length]
}

func maxLength(i, j int) int {
	if i > j {
		return i
	}
	return j
}
