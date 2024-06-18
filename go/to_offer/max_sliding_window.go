package to_offer

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	queue := newQueue(nums, k)
	res := make([]int, len(nums)-k+1)

	left := 0
	right := 0

	for {
		if right >= len(nums) {
			break
		}

		queue.Append(right)
		right++

		if len(nums[left:right]) < k {
			continue
		}

		res[left] = nums[queue.Max()]
		left++
	}

	return res
}

type NodeX struct {
	Index int
	Next  *NodeX
	Pre   *NodeX
}

type Queue struct {
	Head *NodeX
	Tail *NodeX
	k    int
	nums []int
}

func newQueue(nums []int, k int) *Queue {
	return &Queue{
		k:    k,
		nums: nums,
	}
}

func (q *Queue) Append(index int) {
	node := &NodeX{
		Index: index,
	}

	if q.Head == nil {
		q.Head = node
		q.Tail = node
		return
	}

	for {
		if q.Tail == nil {
			q.Head = node
			q.Tail = node
			return
		}

		if q.nums[index] > q.nums[q.Tail.Index] {
			q.Tail = q.Tail.Pre
			if q.Tail != nil {
				q.Tail.Next = nil
			}

			continue
		}

		q.Tail.Next = node
		node.Pre = q.Tail
		q.Tail = node
		break
	}

	return
}

func (q *Queue) Max() int {
	if q.Head == nil {
		return 0
	}

	for {
		if (q.Tail.Index - q.Head.Index) >= q.k {
			q.Head = q.Head.Next
			q.Head.Pre = nil
			continue
		}

		break
	}

	return q.Head.Index
}
