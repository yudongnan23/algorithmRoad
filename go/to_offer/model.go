package to_offer

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type doubleLinkedListNode struct {
	pre   *doubleLinkedListNode
	next  *doubleLinkedListNode
	v     byte
	isDel bool
}

type LinkedList struct {
	head        *doubleLinkedListNode
	tail        *doubleLinkedListNode
	nodeMapping map[byte]*doubleLinkedListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
