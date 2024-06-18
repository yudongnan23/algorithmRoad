package to_offer

func inventoryManagementII(stock []int) int {
	if len(stock) == 0 {
		return 0
	}

	m := make(map[int]int, 0)

	for i := 0; i < len(stock); i++ {
		m[stock[i]] = m[stock[i]] + 1

		if m[stock[i]] > len(stock)/2 {
			return stock[i]
		}
	}

	return 0
}
