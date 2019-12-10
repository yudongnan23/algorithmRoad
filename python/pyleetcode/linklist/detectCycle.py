"""

    给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
    为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
    索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
    说明：不允许修改给定的链表

"""

"""

    思路：定义快慢指针，快指针一次走两步，慢指针一次走一步，当两指针相遇时，将快指针放到起点，与慢指针同时一次走一步，两指针再次相遇时
         即为环的入口
    注：设循环链表实际长度为n，头结点距环的的第一个节点有y步，两指针第一次相遇慢指针走了x步，则快指针走了2x步，则代表此时慢指针离环的
        第一个结点还有n-x步，快指针离环的第一个结点有2x-n步或n-x步，故链表的长度可表示为（2x-n） + （n-x） + y = n得到，y = n-x,
        可知，两指针第一次相遇时，快慢指针距环的第一节点与头结点距环的第一个结点的距离相等。    
    
"""


class Solution:
    def detectCycle(self, head):
        # 定义快慢指针，初始化为链表的头结点
        slow = head
        fast = head
        # 指针往前走的必要条件，否则会抛出异常
        while slow and fast.next and fast.next.next:
            # 慢指针走一步
            slow = slow.next
            # 快指针走两步
            fast = fast.next.next
            # 两指针第一次相遇，将快指针放置起点，即头结点与慢指针同时一次一步往下走直到相遇即为环的入口点
            if slow == fast:
                fast = head
                while slow != fast:
                    slow = slow.next
                    fast = fast.next
                return slow
        return None
