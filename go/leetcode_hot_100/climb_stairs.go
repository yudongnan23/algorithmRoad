package leetcode_hot_100

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	preOne := 2
	preTwo := 1
	ways := preOne
	for i := 2; i < n; i++ {
		ways = preOne + preTwo
		preTwo = preOne
		preOne = ways
	}
	return ways
}
