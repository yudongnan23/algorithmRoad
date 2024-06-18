package leetcode_hot_100

type NodeWithMin struct {
	Val    int
	CurMin int
}

type MinStack struct {
	Stack [][2]int
}

func ConstructorI() MinStack {
	return MinStack{
		Stack: make([][2]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	node := [2]int{val, val}

	if len(m.Stack) > 0 && m.Stack[len(m.Stack)-1][1] < val {
		node[1] = m.Stack[len(m.Stack)-1][1]
	}

	m.Stack = append(m.Stack, node)
}

func (m *MinStack) Pop() {
	m.Stack = m.Stack[:len(m.Stack)-1]
}

func (m *MinStack) Top() int {
	return m.Stack[len(m.Stack)-1][0]
}

func (m *MinStack) GetMin() int {
	return m.Stack[len(m.Stack)-1][1]
}
