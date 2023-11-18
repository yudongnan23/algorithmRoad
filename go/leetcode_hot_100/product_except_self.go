package leetcode_hot_100

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	zeroCount := 0
	n := 1
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		n = n * nums[i]
	}

	for i := 0; i < length; i++ {
		if zeroCount > 1 {
			res[i] = 0
			continue
		}
		if nums[i] == 0 {
			res[i] = n
			continue
		}
		if zeroCount > 0 {
			res[i] = 0
			continue
		}

		res[i] = n / nums[i]
	}

	return res
}
