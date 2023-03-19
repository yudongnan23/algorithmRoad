package to_offer

import (
	"math/big"
)

func fib(n int) {
	if n <= 1 {
		return n
	}

	dynamic := make([]*big.Int, n+1)
	dynamic[0] = big.NewInt(0)
	dynamic[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		adder := big.NewInt(0)
		dynamic[i] = adder.Add(dynamic[i-1], dynamic[i-2])
	}
	modder := big.NewInt(0)
	return int(modder.Mod(dynamic[n], big.NewInt(1000000007)).Int64())
}
