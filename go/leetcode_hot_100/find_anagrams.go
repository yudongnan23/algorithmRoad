package leetcode_hot_100

// TODO Again
func findAnagrams(s string, p string) []int {
	pLen := len(p)
	res := make([]int, 0)

	if pLen > len(s) {
		return res
	}

	newP := sortWords(p)
	subStr := sortWords(s[:pLen])

	if subStr == newP {
		res = append(res, 0)
	}

	n := len(s) - pLen + 1

	for i := 1; i < n; i++ {
		subStr = insertSort(subStr, s[i+pLen-1], s[i-1])
		if subStr == newP {
			res = append(res, i)
		}
	}

	return res
}

func insertSort(src string, newByte, deleteByte byte) string {
	if newByte == deleteByte {
		return src
	}

	searchFunc := func(bs []byte, dst byte) int {
		left := 0
		right := len(bs) - 1
		for left <= right {
			mid := (left + right) / 2
			if bs[mid] == dst {
				left = mid
				break
			}

			if bs[mid] > dst {
				right = mid - 1
				continue
			}

			if bs[mid] < dst {
				left = mid + 1
				continue
			}
		}

		if left >= len(bs) {
			return len(bs) - 1
		}

		return left
	}

	// 删除
	bs := []byte(src)
	deleteIndex := searchFunc(bs, deleteByte)
	newBs := bs[0:deleteIndex]
	newBs = append(newBs, bs[deleteIndex+1:]...)

	// 新增
	b := make([]byte, len(newBs)+1)
	if len(newBs) == 0 {
		b[0] = newByte
		return string(b)
	}

	if newByte <= newBs[0] {
		b[0] = newByte
		copy(b[1:], newBs)
		return string(b)
	}

	if newByte >= newBs[len(newBs)-1] {
		b[len(b)-1] = newByte
		copy(b[:len(b)-1], newBs)
		return string(b)
	}

	newIndex := searchFunc(newBs, newByte)
	copy(b[:newIndex], newBs[:newIndex])
	b[newIndex] = newByte
	copy(b[newIndex+1:], newBs[newIndex:])

	return string(b)
}
