package leetcode_hot_100

// TODO three
func minWindow(s string, t string) string {
	sizeT := len(t)
	sizeS := len(s)
	res := ""
	tMap := make(map[byte]int, 0)
	for i := 0; i < sizeT; i++ {
		tMap[t[i]]++
	}

	sMap := make(map[byte]int, 0)
	left := 0
	right := 0
	for right <= sizeS {
		curSize := right - left
		for curSize >= sizeT {
			if !exists(sMap, tMap) {
				break
			}
			if curSize < len(res) || len(res) == 0 {
				res = s[left:right]
			}
			sMap[s[left]]--
			left++
			curSize = right - left
		}

		if right < sizeS {
			sMap[s[right]]++
		}

		right++
	}

	return res
}

func exists(m1, m2 map[byte]int) bool {
	for k, v := range m2 {
		if m1[k] < v {
			return false
		}
	}
	return true
}
