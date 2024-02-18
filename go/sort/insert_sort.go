package sort

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 1 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
