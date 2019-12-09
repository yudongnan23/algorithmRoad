"""

    编写一个程序，找到两个单链表相交的起始节点

"""

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA, headB):
        point_a = headA
        point_b = headB
        count_a = 0
        count_b = 0
        while point_a or point_b:
            if point_a:
                point_a = point_a.next
            else:
                count_a += 1
            if point_b:
                point_b = point_b.next
            else:
                count_b += 1
        point_a = headA
        point_b = headB
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
        while point_b and point_a:
            if point_b == point_a:
                return point_a
            point_b = point_b.next
            point_a = point_a.next
        return None