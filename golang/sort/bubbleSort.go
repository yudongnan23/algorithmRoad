/*
冒泡排序

    冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
    走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。

    算法描述
        比较相邻的元素。如果第一个比第二个大，就交换它们两个；
        对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
        针对所有的元素重复以上的步骤，除了最后一个；
        重复步骤1~3，直到排序完成。
 */


package main

import "fmt"

func bubbleSort(nums []int) []int{
	length := len(nums)
	if length == 0{
		return nums
	}

	for lastIndex := length - 1; lastIndex > 0; lastIndex--{
		for i := 0; i < lastIndex; i++{
			if nums[i] > nums[i + 1]{
				temp := nums[i + 1]
				nums[i + 1] = nums[i]
				nums[i] = temp
			}
		}
	}
	return nums
}

func main(){
	array := []int{2, 323, 24, 46, 1, 0, 243, 48, 90}
	fmt.Println(bubbleSort(array))
}