package leetcode_hot_100

// TODO three - 跳跃游戏
func canJump(nums []int) bool {
	i, right, n := 0, nums[0], len(nums)-1
	for i <= right && i <= n {
		curRight := i + nums[i]
		if right < curRight {
			right = curRight
		}
		i++
	}
	return right >= n
}
