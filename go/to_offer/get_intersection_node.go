package to_offer

func getIntersectionNode(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var p1Change bool
	var p2Change bool

	p1 := l1
	p2 := l2

	for {
		if p1 == p2 {
			return p1
		}

		p1 = p1.Next
		p2 = p2.Next

		if p1 == nil {
			if p1Change {
				return nil
			}

			p1 = l2
			p1Change = true
		}

		if p2 == nil {
			if p2Change {
				return nil
			}

			p2 = l1
			p2Change = true
		}
	}

	return nil
}
