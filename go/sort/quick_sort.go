package sort

func quickSort(nums []int) {
	var sortFunc func(left, right int)
	sortFunc = func(left, right int) {
		if left >= right {
			return
		}
		nextLeft := left
		nextRight := right
		baseNum := nums[left]
		leftEmpty := true
		rightEmpty := false
		for left < right {
			if baseNum > nums[right] && leftEmpty {
				nums[left] = nums[right]
				left++
				leftEmpty = false
				rightEmpty = true
				continue
			}

			if baseNum < nums[left] && rightEmpty {
				nums[right] = nums[left]
				leftEmpty = true
				rightEmpty = false
				right--
				continue
			}

			if rightEmpty {
				left++
			} else {
				right--
			}
		}

		nums[left] = baseNum
		sortFunc(nextLeft, left-1)
		sortFunc(left+1, nextRight)
	}

	sortFunc(0, len(nums)-1)
}
