package to_offer

type MSNode struct {
	Val    int
	Next   *MSNode
	curMin int
}

type MinStack struct {
	Head *MSNode
}

func MSConstructor() MinStack {
	return MinStack{
		Head: nil,
	}
}

func (ms *MinStack) Push(x int) {
	node := &MSNode{
		Val:    x,
		curMin: x,
	}

	if ms.Head == nil {
		ms.Head = node
		return
	}

	if ms.Head.curMin < x {
		node.curMin = ms.Head.curMin
	}

	node.Next = ms.Head
	ms.Head = node
	return
}

func (ms *MinStack) Pop() {
	if ms.Head == nil {
		return
	}

	ms.Head = ms.Head.Next
	return
}

func (ms *MinStack) Top() int {
	if ms.Head == nil {
		return 0
	}

	return ms.Head.Val
}

func (ms *MinStack) Min() int {
	if ms.Head == nil {
		return 0
	}

	return ms.Head.curMin
}
