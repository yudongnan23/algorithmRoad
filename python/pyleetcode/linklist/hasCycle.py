"""

    给定一个链表，判断链表中是否有环。

    为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
    （索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

"""
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head):
        result = False
        if head is None:
            return result
        slow = head
        quick = head
        while slow.next and quick.next and quick.next.next:
            slow = slow.next
            quick = quick.next.next
            if slow == quick:
                result = True
                break
        return result
