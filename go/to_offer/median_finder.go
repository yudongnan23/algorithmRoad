package to_offer

// TODO 待使用堆排序优化

type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func ConstructorI() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	nums := make([]int, len(mf.nums)+1)
	i := 0
	j := 0
	for i < len(mf.nums) {
		if mf.nums[i] > num && j == i {
			nums[j] = num
			copy(nums[j+1:], mf.nums[i:])
			i = len(mf.nums)
			j = i + 1
			break
		}
		nums[j] = mf.nums[i]
		j++
		i++
	}

	if j == i {
		nums[j] = num
	}

	mf.nums = nums
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
