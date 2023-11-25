package leetcode_hot_100

// Node
// TODO again
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var (
		newHead *Node
		newP    *Node
	)
	p := head
	new2OldRandomMap := make(map[*Node]*Node, 0)
	old2NewMap := make(map[*Node]*Node, 0)
	for p != nil {
		node := &Node{
			Val: p.Val,
		}
		old2NewMap[p] = node
		if p.Random != nil {
			new2OldRandomMap[node] = p.Random
		}
		p = p.Next
		if newHead == nil {
			newHead = node
			newP = node
			continue
		}

		newP.Next = node
		newP = newP.Next
	}

	newP = newHead
	for newP != nil {
		oldRandom, ok := new2OldRandomMap[newP]
		if ok {
			newP.Random = old2NewMap[oldRandom]
		}

		newP = newP.Next
	}

	return newHead
}
