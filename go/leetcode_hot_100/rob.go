package leetcode_hot_100

func rob(nums []int) int {
	length := len(nums)
	newNums := make([]int, length+3)
	copy(newNums[3:], nums)
	for i := 3; i < length+3; i++ {
		newNums[i] = max(newNums[i-3]+newNums[i], newNums[i-2]+newNums[i])
	}
	return max(newNums[len(newNums)-1], newNums[len(newNums)-2])
}
