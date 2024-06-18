package leetcode_hot_100

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) == 0 || k == 0 {
		return
	}

	n := make([]int, k)
	copy(n, nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], n)
}
