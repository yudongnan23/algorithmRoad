package to_offer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newRandom2OldMap := make(map[*Node]*Node, 0)
	old2NewMap := make(map[*Node]*Node, 0)
	var newHead *Node

	newP := newHead
	p := head
	for {
		if p == nil {
			break
		}

		newNode := &Node{
			Val: p.Val,
		}

		old2NewMap[p] = newNode
		if p.Random != nil {
			newRandom2OldMap[newNode] = p.Random
		}

		p = p.Next
		if newHead == nil {
			newHead = newNode
			newP = newNode
			continue
		}

		newP.Next = newNode
		newP = newP.Next
	}

	newP = newHead

	for {
		if newP == nil {
			break
		}

		old, ok := newRandom2OldMap[newP]
		if !ok {
			newP = newP.Next
			continue
		}

		newRandom := old2NewMap[old]
		newP.Random = newRandom
		newP = newP.Next
	}

	return newHead
}
