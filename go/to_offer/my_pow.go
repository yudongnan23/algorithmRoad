package to_offer

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}

	return x * temp * temp
}
