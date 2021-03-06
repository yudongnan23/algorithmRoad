/*
	给定指定的硬币种类及需要的钱的总数，组合给定的硬币得到钱的总数，并使得硬币数最少。
*/

package main

import "fmt"

func coinExchange(coinArray []int, moneny int) int {
	const MAXVALUE = 9999
	length := len(coinArray)
	f := make([]int, moneny+1)
	f[0] = 0
	for i := 1; i <= moneny; i++ {
		minNum := MAXVALUE
		for j := 0; j < length; j++ {
			if i >= coinArray[j] && minNum > f[i-coinArray[j]]+1 {
				minNum = f[i-coinArray[j]] + 1
			}
		}
		f[i] = minNum
	}

	if f[moneny] == MAXVALUE {
		return -1
	}
	return f[moneny]
}

func main() {
	coin := []int{2, 5, 7}
	money := 27
	fmt.Println(coinExchange(coin, money))
}
