"""

    请判断一个链表是否为回文链表。

"""


"""

    思路：定义一个函数使用快慢指针找到链表的中点，再定义一个函数将中点后面的结点进行反转，并返回反转后的头结点
            对链表进行双向同时遍历按，当遇两个节点的节点值不相等时，代表不是该链表不是回文链表，否则是回文链表。

"""


class Solution:
    def isPalindrome(self, head):
        # 调用函数获取链表中点
        middle_node = self.get_middle_node(head)
        # 调用反转函数对中点后面的结点进行反转并返回反转后的头结点
        new_head = self.reverse(middle_node)
        # 左指针
        left_point = head
        # 右指针
        right_point = new_head
        result = True
        # 让两个指针同时走，每个指针一次走一步
        while left_point and right_point:
            # 当两指针的结点值不相等时，代表链表不是回文链表
            if left_point.val != right_point.val:
                result = False
                break
            left_point = left_point.next
            right_point = right_point.next
        return result

    def reverse(self, middle):
        if middle is None:
            return middle
        # 递归函数出口，即递归到链表尾结点
        if middle.next is None:
            return middle
        # 递归调用本身对链表的每个节点进行反转
        new_head = self.reverse(middle.next)
        # 当前结点与next结点进行反转
        cur = middle
        middle.next.next = cur
        middle.next = None
        return new_head

    def get_middle_node(self, head):
        # 定义慢指针，一次走一步
        slow = head
        # 定义块指针，一次走两步
        fast = head
        # 当快指针走到链表尾部时，慢指针刚好在中点位置
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        return slow
