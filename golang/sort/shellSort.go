/*
希尔排序
    1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。

    算法描述
        先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：

            选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
            按增量序列个数k，对序列进行k 趟排序；
            每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
 */

package main

import "fmt"

func shellSort(nums []int)[]int{
	length := len(nums)

	if length == 0{
		return nums
	}

	for gap := length / 2; gap > 0; gap = gap / 2{
		for curIndex := gap; curIndex < length; curIndex ++{
			variantIndex := curIndex
			temp := nums[variantIndex]

			for ; variantIndex - gap >= 0 ;{
				if temp < nums[variantIndex - gap]{
					nums[variantIndex] = nums[variantIndex - gap]
					variantIndex = variantIndex - gap
				}else{
					break
				}
			}
			nums[variantIndex] = temp
		}
	}

	return nums
}

func main(){
	array := []int{1, 2, 1, 8, 7, 19, 0}
	fmt.Println(shellSort(array))
}