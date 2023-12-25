package leetcode_hot_100

func jump(nums []int) int {
	count, end, farthest := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		curFarthest := i + nums[i]
		if curFarthest > farthest {
			farthest = curFarthest
		}
		if i == end {
			end = farthest
			count++
		}
	}
	return count
}
