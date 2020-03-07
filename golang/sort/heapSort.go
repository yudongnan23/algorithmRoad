/*
堆排序

    堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，
	并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

    算法描述
        将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
        将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
        由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，
		得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成


*/

package main

import (
	"fmt"
)

func heapSort(nums []int) []int{
	length := len(nums)

	if length == 0 || length == 1{
		return nums
	}

	buildHeap(nums, length)

	for curNode := length - 1; curNode > 0; curNode--{
		Swap(nums, 0, curNode)
		heapify(nums, curNode, 0)
	}

	return nums
}

func buildHeap(tree []int, length int){
	lastNode := length - 1
	parentNode := (lastNode - 1) / 2

	for curNode := parentNode; curNode >= 0; curNode--{
		heapify(tree, length, curNode)
	}
}

func heapify(tree []int, length int, curNode int){
	if curNode >= length{
		return
	}

	s1 := curNode * 2 + 1
	s2 := curNode * 2 + 2
	maxNode := curNode

	if s1 < length && tree[s1] > tree[maxNode]{
		maxNode = s1
	}

	if s2 < length && tree[s2] > tree[maxNode]{
		maxNode = s2
	}

	if maxNode != curNode{
		Swap(tree, curNode, maxNode)
		heapify(tree, length, maxNode)
	}
}

func Swap(nums []int, index1 int, index2 int){
	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}

func main(){
	array := []int{1, 0, 2, 8, 0, 100, 2, 0, 100}
	fmt.Println(heapSort(array))
}