/*
快速排序

    算法描述
        快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

    算法描述
        快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

            从数列中挑出一个元素，称为 “基准”（pivot）；
            重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
            递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
 */

package main

import "fmt"

func quickSort(nums []int)[]int{
	length := len(nums)

	if length == 0{
		return nums
	}

	left := 0
	right := length - 1

	sort(nums, left, right)

	return nums
}

func sort(nums []int, left int, right int){
	if left < right{
		partitionIndex := getPartitionIndex(nums, left, right)

		sort(nums, left, partitionIndex - 1)
		sort(nums, partitionIndex + 1, right)
	}
}

func getPartitionIndex(nums []int, left int, right int)int{
	pivot := left
	curIndex := pivot + 1

	for variantIndex := curIndex; variantIndex < right + 1; variantIndex ++{
		if nums[variantIndex] < nums[pivot]{
			swap(nums, curIndex, variantIndex)
			curIndex ++
		}
	}

	swap(nums, pivot, curIndex - 1)

	return curIndex - 1
}

func swap(nums []int, i int, j int){
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}

func main(){
	array := []int{1, 4, 5, 0, 10, 0, 1, 100, 9, 20}
	fmt.Println(quickSort(array))
}