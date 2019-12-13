"""

    给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

"""


"""

    思路：典型的双指针解法，新建一个链表结点node，将其next指向head头结点，定义两个指针，left和right，left指针初始化node，
        right指针初始化head头指针；让right指针先走n步，然后两指针同时走，直到right指针到达链表尾部，此时left指针刚好在
        倒数第n个结点的前一个节点，此时便可将倒数第n个结点删除。

"""


class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None


class Solution:
    def removeNthFromEnd(self, head, n):
        # 新建一个结点，并将其next指向头结点
        node = ListNode(-1)
        node.next = head
        # 定义left指针，初始化为node
        left = node
        # 定义right指针，初始化为head
        right = head
        count = 0
        # 让right指针先走n步，注意倒数没有0
        while count < n - 1:
            right = right.next
            count += 1
        # 两指针同时走，直到right到达链表尾部
        while right.next:
            left = left.next
            right = right.next
        # 如果指针没动，即，代表要删除头结点head，将head指向head.next即可
        if left == node:
            head = head.next
        else:
            # 执行结点删除
            left.next = left.next.next
        return head
