"""

    给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
    为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
    索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
    说明：不允许修改给定的链表

"""
class Solution:
    def detectCycle(self, head):
        slow = head
        fast = head
        while slow and fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                fast = head
                while slow != fast:
                    slow = slow.next
                    fast = fast.next
                return slow
        return None
