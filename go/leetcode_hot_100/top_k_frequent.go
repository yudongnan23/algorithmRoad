package leetcode_hot_100

type ValWithCount struct {
	Val   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	d := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		d[nums[i]]++
	}

	newNums := make([]ValWithCount, 0, len(d))
	for k, v := range d {
		newNums = append(newNums, ValWithCount{
			Val:   k,
			Count: v,
		})
	}

	newNums = mergeSortII(newNums)

	res := make([]int, 0, k)
	for i := len(newNums) - 1; len(res) < k; i-- {
		res = append(res, newNums[i].Val)
	}
	return res
}

func mergeSortII(nums []ValWithCount) []ValWithCount {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	nums1 := mergeSortII(nums[:mid])
	nums2 := mergeSortII(nums[mid:])
	return sortTwoSliceII(nums1, nums2)
}

func sortTwoSliceII(nums1, nums2 []ValWithCount) []ValWithCount {
	nums := make([]ValWithCount, 0, len(nums1)+len(nums2))
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

		if nums1[index1].Count < nums2[index2].Count {
			nums = append(nums, nums1[index1])
			index1++
		} else {
			nums = append(nums, nums2[index2])
			index2++
		}
	}

	return nums
}
