package sort

import (
	"fmt"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 8, 6, 5, 7, 10, 9}
	bubbleSort(nums)
	fmt.Println(nums)
}

func Test_selectSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 8, 6, 5, 7, 10, 9}
	selectSort(nums)
	fmt.Println(nums)
}
