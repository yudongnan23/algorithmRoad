package sort

import (
	"fmt"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	bubbleSort(nums)
	fmt.Println(nums)
}

func Test_selectSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 4, 3, 7, 8, 6, 5}
	selectSort(nums)
	fmt.Println(nums)
	nums = []int{2, 2, 2, 3, 3, 3, 1, 1, 1, 1}
	selectSort(nums)
	fmt.Println(nums)
}

func Test_insertSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	insertSort(nums)
	fmt.Println(nums)
}

func Test_shellSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	shellSort(nums)
	fmt.Println(nums)
}

func Test_quickSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	quickSort(nums)
	fmt.Println(nums)
}

func Test_mergeSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	mergeSort(nums)
	fmt.Println(nums)
}

func Test_countSort(t *testing.T) {
	nums := []int{9, 10, 1, 2, 5, 6, 8, 4, 7, 3}
	countSort(nums)
	fmt.Println(nums)
}

func Test_heapSort(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	heapSort(nums)
	fmt.Println(nums)
}
