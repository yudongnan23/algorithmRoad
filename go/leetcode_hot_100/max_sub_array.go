package leetcode_hot_100

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
