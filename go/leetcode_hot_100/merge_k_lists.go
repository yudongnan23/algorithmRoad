package leetcode_hot_100

// TODO three - 合并 K 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}
