/*
归并排序
    归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
    将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

    算法描述
        把长度为n的输入序列分成两个长度为n/2的子序列；
        对这两个子序列分别采用归并排序；
        将两个排序好的子序列合并成一个最终的排序序列。
 */

package main

import "fmt"

func mergeSort(nums []int)[]int{
	length := len(nums)

	if length == 0 || length == 1{
		return nums
	}

	middle := length / 2

	return merge(mergeSort(nums[0: middle]), mergeSort(nums[middle: length]))
}

func merge(leftNums []int, rightNums []int)[]int{
	newNums := []int{}

	leftNumsLength := len(leftNums)
	rightNumsLength := len(rightNums)

	leftIndex := 0
	rightIndex := 0

	for ; leftIndex < leftNumsLength || rightIndex < rightNumsLength; {
		if leftIndex < leftNumsLength && rightIndex < rightNumsLength{
			if leftNums[leftIndex] < rightNums[rightIndex]{
				newNums = append(newNums, leftNums[leftIndex])
				leftIndex ++
			}else{
				newNums = append(newNums, rightNums[rightIndex])
				rightIndex ++
			}
		}else if leftIndex < leftNumsLength{
			newNums = append(newNums, leftNums[leftIndex: leftNumsLength]...)
			break
		}else if rightIndex < rightNumsLength{
			newNums = append(newNums, rightNums[rightIndex: rightNumsLength]...)
			break
		}
	}

	return newNums
}

func main(){
	array := []int{10, 1, 0, 19, 100, 200, 0, 1, 27}
	//nums1 := []int{1, 2, 3, 6}
	//nums2 := []int{1, 2, 4, 6, 7, 8}
	fmt.Println(mergeSort(array))
}