"""

    删除链表中等于给定值 val 的所有节点。


"""


"""

    思路： 新建一结点，将其next指向链表头结点，初始化一个指针为该新建的结点，然后让指针往下走，当其Next结点的结点值等于目标值时，
         将其Next结点删除即可，最后返回新建结点的Next结点即可。

"""


class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None


class Solution:
    def removeElements(self, head, val):
        # 新建一个结点
        node = ListNode(-1)
        node.next = head
        # 初始化指针为新建的结点
        point = node
        # 让指针往下走，直到指针的Next指针为空，代表指针已走到链表尾部
        while point.next:
            # 当指针Next结点的节点值为目标值，执行删除删除操作
            if point.next.val == val:
                point.next = point.next.next
            # 否则指针往前走一步
            else:
                point = point.next
        # 得到链表新的头结点并返回
        new_head = node.next
        return new_head
