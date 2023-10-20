package to_offer

import (
	"math/big"
)

var _lenII = map[int]int{
	0: 1,
	1: 1,
	2: 1,
	3: 2,
	4: 4,
	5: 6,
	6: 9,
}

func cuttingBambooII(length int) int {
	if length <= 6 {
		return _len[length]
	}

	dp := make([]*big.Int, length+1)
	for i := 0; i <= 6; i++ {
		dp[i] = big.NewInt(int64(_len[i]))
	}

	for i := 7; i <= length; i++ {
		x := big.NewInt(1)
		dp[i] = maxII(x.Mul(dp[i-2], big.NewInt(2)), x.Mul(dp[i-3], big.NewInt(3)))
	}

	return int(dp[length].Mod(dp[length], big.NewInt(1000000007)).Int64())
}

func maxII(i, j *big.Int) *big.Int {
	if i.Cmp(j) == 1 {
		return i
	}
	return j
}
