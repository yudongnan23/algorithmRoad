package to_offer

// TODO 待使用堆排序优化

type MedianFinder struct {
	nums       []int
	backupNums []int
}

func ConstructorXX() MedianFinder {
	return MedianFinder{
		nums:       make([]int, 0),
		backupNums: make([]int, 0),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	size := len(mf.nums)
	mf.backupNums = append(mf.backupNums, 0)
	if size > 0 {
		mf.backupNums = append(mf.backupNums, 0)
	}

	defer func() {
		mf.nums, mf.backupNums = mf.backupNums, mf.nums
	}()

	if size == 0 || mf.nums[size-1] <= num {
		copy(mf.backupNums[:size], mf.nums)
		mf.backupNums[size] = num
		return
	}

	if mf.nums[0] >= num {
		copy(mf.backupNums[1:size+1], mf.nums)
		mf.backupNums[0] = num
		return
	}

	left := 0
	right := len(mf.nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if mf.nums[mid] == num {
			left = mid
			break
		}

		if mf.nums[mid] > num {
			right = mid - 1
			continue
		}

		if mf.nums[mid] < num {
			left = mid + 1
			continue
		}
	}

	if mf.nums[left] > num {
		left--
	}

	copy(mf.backupNums[:left+1], mf.nums[:left+1])
	mf.backupNums[left+1] = num
	copy(mf.backupNums[left+2:], mf.nums[left+1:])
}

func (mf *MedianFinder) FindMedian() float64 {
	length := len(mf.nums)
	if length == 0 {
		return 0
	}
	mid := length / 2
	if length%2 == 1 {
		return float64(mf.nums[mid])
	}

	return float64(mf.nums[mid]+mf.nums[mid-1]) / float64(2)
}
