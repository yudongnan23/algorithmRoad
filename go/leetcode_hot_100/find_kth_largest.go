package leetcode_hot_100

// TODO again
func findKthLargest(nums []int, k int) int {
	newNums := mergeSort(nums)
	return newNums[len(newNums)-k]
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	nums1 := mergeSort(nums[:mid])
	nums2 := mergeSort(nums[mid:])
	return sortTwoSlice(nums1, nums2)
}

func sortTwoSlice(nums1, nums2 []int) []int {
	nums := make([]int, 0, len(nums1)+len(nums2))
	index1 := 0
	index2 := 0

	for index1 < len(nums1) || index2 < len(nums2) {
		if index1 >= len(nums1) {
			nums = append(nums, nums2[index2:]...)
			break
		}

		if index2 >= len(nums2) {
			nums = append(nums, nums1[index1:]...)
			break
		}

		if nums1[index1] < nums2[index2] {
			nums = append(nums, nums1[index1])
			index1++
		} else {
			nums = append(nums, nums2[index2])
			index2++
		}
	}

	return nums
}
