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
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	selectSort(nums)
	fmt.Println(nums)
}

func Test_insertSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	insertSort(nums)
	fmt.Println(nums)
}

func Test_shellSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 8, 6, 5, 7, 10, 9}
	shellSort(nums)
	fmt.Println(nums)
}

func Test_quickSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	quickSort(nums)
	fmt.Println(nums)
}

func Test_mergeSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	mergeSort(nums)
	fmt.Println(nums)
}

func Test_countSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	countSort(nums)
	fmt.Println(nums)
}
