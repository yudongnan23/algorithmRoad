package to_offer

func maxSales(sales []int) int {
	if len(sales) == 0 {
		return 0
	}

	maxSale := sales[0]
	saleSum := sales[0]

	for i := 1; i < len(sales); i++ {
		if sales[i]+saleSum > sales[i] {
			saleSum = sales[i] + saleSum
		} else {
			saleSum = sales[i]
		}

		if saleSum > maxSale {
			maxSale = saleSum
		}
	}

	return maxSale
}
