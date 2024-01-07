package leetcode_hot_100

// TODO three
func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	res[0] = 1
	for i := 1; i < size; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for i := size - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r = r * nums[i]
	}

	return res
}
