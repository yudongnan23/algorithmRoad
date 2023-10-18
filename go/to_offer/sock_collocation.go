package to_offer

func sockCollocation(sockets []int) []int {
	m := make(map[int]int, 0)

	for i := 0; i < len(sockets); i++ {
		m[sockets[i]] = m[sockets[i]] + 1
	}

	res := make([]int, 0)
	for k, count := range m {
		if count == 1 {
			res = append(res, k)
		}
		if len(res) == 2 {
			break
		}
	}

	return res
}
