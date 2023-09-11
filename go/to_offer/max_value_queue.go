package to_offer

type Value struct {
	Val     int
	nodePtr *NodeXX
}

type NodeXX struct {
	Val  int
	Next *NodeXX
	Pre  *NodeXX
}

type MaxQueue struct {
	Vals  []Value
	Front *NodeXX
	Back  *NodeXX
}

func ConstructorX() MaxQueue {
	return MaxQueue{
		Vals:  make([]Value, 0),
		Back:  nil,
		Front: nil,
	}
}

func (q *MaxQueue) Max_value() int {
	if q.Front == nil {
		return -1
	}

	return q.Front.Val
}

func (q *MaxQueue) Push_back(val int) {
	node := &NodeXX{
		Val: val,
	}

	q.Vals = append(q.Vals, Value{
		Val:     val,
		nodePtr: node,
	})

	if q.Front == nil {
		q.Front = node
		q.Back = node
		return
	}

	for {
		if q.Back == nil {
			q.Back = node
			q.Front = node
			break
		}

		if q.Back.Val >= val {
			q.Back.Next = node
			node.Pre = q.Back
			q.Back = node
			return
		}

		q.Back = q.Back.Pre
		if q.Back != nil {
			q.Back.Next = nil
		}
	}

	return
}

func (q *MaxQueue) Pop_front() int {
	if len(q.Vals) == 0 {
		return -1
	}

	popVal := q.Vals[0]
	q.Vals = q.Vals[1:]

	if q.Front != nil && popVal.nodePtr == q.Front {
		q.Front = q.Front.Next
		if q.Front != nil {
			q.Front.Pre = nil
		}
	}

	return popVal.Val
}
