package sort

func shellSort(nums []int) {
	length := len(nums)
	gap := length / 2
	for gap > 0 {
		for i := 0; i < length; i = i + gap {
			for j := i + gap; j < length && j >= gap && nums[j] < nums[j-gap]; j = j - gap {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
			}
		}

		gap = gap / 2
	}
}
