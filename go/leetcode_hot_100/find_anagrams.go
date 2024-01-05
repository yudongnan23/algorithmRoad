package leetcode_hot_100

// TODO three
func findAnagrams(s string, p string) []int {
	targetSize := len(p)
	pCount := [26]int{}
	for i := 0; i < targetSize; i++ {
		pCount[int(p[i])-97]++
	}

	size := len(s)
	l := 0
	r := 0
	res := make([]int, 0, size)
	slideCount := [26]int{}
	for r < size {
		slideCount[int(s[r])-97]++
		if r-l < targetSize-1 {
			r++
			continue
		}

		if isAnagram(slideCount, pCount) {
			res = append(res, l)
		}
		slideCount[int(s[l])-97]--
		r++
		l++
	}
	return res
}

func isAnagram(m1, m2 [26]int) bool {
	for i := 0; i < len(m1); i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
