package leetcode_hot_100

// TODO again
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	d := make(map[string]int, 0)
	res := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		index, ok := d[sortWords(strs[i])]
		if !ok {
			res = append(res, []string{strs[i]})
			d[sortWords(strs[i])] = len(res) - 1
			continue
		}

		res[index] = append(res[index], strs[i])
	}

	return res
}

func sortWords(word string) string {
	bs := []byte(word)
	for i := 0; i < len(bs); i++ {
		for j := 0; j < i; j++ {
			if bs[j] <= bs[i] {
				bs[j], bs[i] = bs[i], bs[j]
			}
		}
	}
	return string(bs)
}
