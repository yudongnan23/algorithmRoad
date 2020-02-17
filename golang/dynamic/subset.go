/*
	给定一个数组arr，和一个目标数S，从数组随机挑出任意个数，若能使得这几个数的和等于目标数S，返回True，否则返回False
*/

package main

import "fmt"

// recursive递归解法
func recursiveSubset(arr []int, curIndex int, S int) bool {
	if S == 0 || arr[curIndex] == S {
		return true
	} else if curIndex == 0 && arr[curIndex] != S {
		return false
	} else if arr[curIndex] > S {
		return recursiveSubset(arr, curIndex-1, S)
	} else {
		selectSubset := recursiveSubset(arr, curIndex-1, S-arr[curIndex])
		notSelectSubset := recursiveSubset(arr, curIndex-1, S)
		return selectSubset || notSelectSubset
	}
}

// dynamic programing动态规划解法
func dp_subset(arr []int, S int) bool {
	arrLength := len(arr)
	subset := make([][]bool, len(arr))
	for i := 0; i < arrLength; i++ {
		sub := make([]bool, S+1)
		sub[0] = true
		if i == 0 {
			sub[arr[0]] = true
		}
		subset[i] = sub
	}
	for i := 1; i < arrLength; i++ {
		for s := 1; s < S+1; s++ {
			if arr[i] > s {
				subset[i][s] = subset[i-1][s]
			} else {
				selectSubset := subset[i-1][s-arr[i]]
				notSelectSubset := subset[i-1][s]
				subset[i][s] = selectSubset || notSelectSubset
			}
		}
	}
	return subset[arrLength-1][S]
}

func main() {
	arr := []int{3, 34, 4, 12, 5, 2}
	index := 5
	S := 9
	fmt.Println(recursiveSubset(arr, index, S))
	fmt.Println(dp_subset(arr, S))
}
