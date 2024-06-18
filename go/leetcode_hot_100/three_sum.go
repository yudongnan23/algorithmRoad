package leetcode_hot_100

import "sort"

// TODO three - 三数之和
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				leftVal := nums[l]
				for nums[l] == leftVal && l < r {
					l++
				}

				rightVal := nums[r]
				for nums[r] == rightVal && r > l {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
