"""

    编写一个程序，找到两个单链表相交的起始节点

"""

"""

    思路：求出两个链表的长度差，然后让指向长的的链表的指针先走长度差个单位，然后指向两个链表的指针同时，当两个指针相等时，代表两个链表的相交点。

"""

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA, headB):
        # 定义两个指针，初始化指向两个链表的表头
        point_a = headA
        point_b = headB
        # 定义两个常量，初始化为空，用于记录长度差
        count_a = 0
        count_b = 0
        # 对两个两个链表进行遍历，遍历条件为两个链表不同时为空，当其中一个为空时，另外一个链表继续遍历，由count_a或者count_b记录长度差，分别对应的链表
        while point_a or point_b:
            # 指针若不为空，继续往下走，否则，将长度差常量进行加1操作，下同。
            if point_a:
                point_a = point_a.next
            else:
                count_a += 1
            if point_b:
                point_b = point_b.next
            else:
                count_b += 1
        # 重新定义指针，初始化为两个链表的表头
        point_a = headA
        point_b = headB
        # 判断记录长度差的常量是否为0，若不为0则代表相应的链表长度较短，对指向另外一个链表表头的指针进行走长度差个单位操作
        if count_b != 0:
            count = 0
            while count < count_b:
                point_a = point_a.next
                count += 1
        elif count_a != 0:
            count = 0
            while count < count_a:
                point_b = point_b.next
                count += 1
        # 此时两指针已经相对同步，可让其同时往下走，直到两个指针相等。
        while point_b and point_a:
            if point_b == point_a:
                return point_a
            point_b = point_b.next
            point_a = point_a.next
        return None