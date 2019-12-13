"""

    反转一个单链表。

"""


"""

    思路：
        1，递归：
            将函数定义为递归函数，递归出口为当前结点为None或head.next为None时，将head返回即可，递归函数调用递归函数进行下一结点的
            反转，入参当前结点的next结点，然后再进行当前结点与next结点的反转。
        2，迭代：
            定义两个指针，第一个指针指向新链表的头结点，初始化为head结点，第二个指针指向待反转链表的头结点，初始化为head.next结点
            然后对链表进行遍历，将待反转的链表表头一个一个放到新链表的表头，直到待反转链表无节点。

"""

class Solution:
    # 递归解法
    def reverseList(self, head):
        # 空判断，当head为None，返回head
        if head is None:
            return head
        # 递归出口，当head.next为None，代表遍历已经到达链表尾部，返回当前结点即可
        if head.next is None:
            return head
        # 将新链表头结点始终指向下一结点的递归调用，最后遇到递归出口时，返回的是链表尾结点，而在每个节点的递归调用过程中，会对每个结点与其next结点进行反转
        new_head = self.reverseList(head.next)
        # 当前结点与next结点进行反装
        cur = head
        cur.next.next = head
        # 当前结点的的next指向为None，此步只在最后一步起作用
        head.next = None
        # 返回新链表头结点
        return new_head

    # 迭代解法
    def reverselist(self, head):
        # 空判断
        if head is None:
            return head
        # 定义指针，该指针始终指向新链表表头
        left = head
        # 定义指针，该指针始终指向待反转链表的头结点
        right = head.next
        # 头结点的next需要指向None，因为反转后head结点将会是新链表的尾结点
        head.next = None
        # 对待反转链表进行遍历，直到指向其头结点的指针为空
        while right:
            cur = right
            #  指向待反转链表头结点的指针指向下一个结点
            right = right.next
            # 将待反转链表的头结点放在新链表的头部
            cur.next = left
            # 新链表头结点指针指向新的头结点
            left = cur
        return left
