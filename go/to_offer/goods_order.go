package to_offer

func goodsOrder(goods string) []string {
	if len(goods) == 0 {
		return nil
	}

	duplicate := make(map[string]bool, 0)

	res := make([]string, 0)

	dfsIIII(&res, duplicate, goods, len(goods), "")

	return res
}

func dfsIIII(res *[]string, duplicate map[string]bool, goods string, length int, current string) {
	if len(current) == length {
		if duplicate[current] {
			return
		}
		duplicate[current] = true
		*res = append(*res, current)
		return
	}

	for i := 0; i < len(goods); i++ {
		dfsIIII(res, duplicate, goods[0:i]+goods[i+1:], length, current+goods[i:i+1])
	}
}
