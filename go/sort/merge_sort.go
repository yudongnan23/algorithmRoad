package sort

func mergeSort(nums []int) {
	var sortFunc func(left, mid, right int)
	sortFunc = func(left, mid, right int) {
		if left >= right {
			return
		}

		if right-left > 1 {
			sortFunc(left, (left+mid)/2, mid-1)
			sortFunc(mid, (mid+right)/2, right)
		}

		newNums := make([]int, 0, right-left+1)
		i := left
		j := mid
		for i < mid && j <= right {
			if nums[i] < nums[j] {
				newNums = append(newNums, nums[i])
				i++
				continue
			}
			newNums = append(newNums, nums[j])
			j++
		}

		if i == mid {
			newNums = append(newNums, nums[j:right+1]...)
		} else {
			newNums = append(newNums, nums[i:mid]...)
		}

		copy(nums[left:right+1], newNums)
	}

	sortFunc(0, len(nums)/2, len(nums)-1)
}
