package leetcode_hot_100

// TODO again
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

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

	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2]) + float64(nums[(len(nums)/2)-1])) / 2
	}

	return float64(nums[len(nums)/2])
}
