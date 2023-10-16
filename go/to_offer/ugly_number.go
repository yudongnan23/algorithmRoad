package to_offer

// TODO again

func nthUglyNumber(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1

	xi := 0
	yi := 0
	zi := 0

	for i := 1; i < n; i++ {
		x := dp[xi] * 2
		y := dp[yi] * 3
		z := dp[zi] * 5

		dp[i] = min(x, y, z)

		if dp[i] == x {
			xi++
		}

		if dp[i] == y {
			yi++
		}

		if dp[i] == z {
			zi++
		}
	}

	return dp[n-1]
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}

	if y <= x && y <= z {
		return y
	}

	return z
}
