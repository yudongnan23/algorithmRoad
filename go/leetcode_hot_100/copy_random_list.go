package leetcode_hot_100

// Node
// TODO three
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var (
		newHead          = &Node{}
		newP             = newHead
		new2OldRandomMap = make(map[*Node]*Node, 0)
		old2NewMap       = make(map[*Node]*Node, 0)
	)

	for p := head; p != nil; p = p.Next {
		node := &Node{Val: p.Val}
		newP.Next = node
		newP = newP.Next
		new2OldRandomMap[node] = p.Random
		old2NewMap[p] = node
	}

	for p := newHead.Next; p != nil; p = p.Next {
		if new2OldRandomMap[p] == nil {
			continue
		}
		p.Random = old2NewMap[new2OldRandomMap[p]]
	}

	return newHead.Next
}
