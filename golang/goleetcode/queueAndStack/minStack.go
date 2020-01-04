/*
	    设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

			push(x) -- 将元素 x 推入栈中。
			pop() -- 删除栈顶的元素。
			top() -- 获取栈顶元素。
			getMin() -- 检索栈中的最小元素。
 */


package main

import (
	"fmt"
)

type MinStack struct {
	stack []int
	top *int
	min *int
}

func Constructor() MinStack {
	minStack := MinStack{}
	return minStack
}

func (this *MinStack)Push(x int){
	this.stack = append(this.stack, x)
	this.top = &x
	if this.min == nil || x < *this.min{
		this.min = &x
	}
}

func (this *MinStack)Pop(){
	length := len(this.stack)
	if length != 0{
		pop := this.stack[length-1]
		this.stack = this.stack[0:length-1]
		newLength := len(this.stack)
		if newLength != 0{
			this.top = &this.stack[newLength-1]
			if pop == *this.min{
				*this.min = getMin(this.stack)
			}
		}else{
			this.min = nil
			this.top = nil
		}

	}
}

func getMin(stack []int) int{
	minNum := stack[0]
	for _, value := range stack{
		if value < minNum{
			minNum = value
		}
	}
	return minNum
}

func (this *MinStack)Top()int{
	return *this.top
}

func (this *MinStack)GetMin()int{
	return *this.min
}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.Pop()
	fmt.Print(stack.GetMin())
}