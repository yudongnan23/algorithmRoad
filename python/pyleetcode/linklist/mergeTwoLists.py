"""

    将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

"""


"""

    思路： 定义四个指针，指针point1，用于遍历l1链表，指针point2，用于遍历l2，指针point，用于指向新链表的尾结点，指针newHead用于指向新链表头结点
            对两个链表进行遍历，将两指针指向的结点结点值较小的一个结点构造成链表结点连接至新链表尾结点，然后重新进行下一轮遍历与比较。

"""


class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None


class Solution:
    def mergeTwoLists(self, l1, l2):
        point1 = l1
        point2 = l2
        new_head = None
        point = None
        # 当两指针其一不为None时
        while point1 or point2:
            # 当两指针都不为None时，代表point1与point2指向的结点都有值
            if point1 and point2:
                val1 = point1.val
                val2 = point2.val
                # 当val1小于val2时，将val1构造成结点放置新链表尾部
                if val1 <= val2:
                    node = ListNode(val1)
                    # 判断新链表头结点是否为None，若是，则将新链表头结点与尾结点指针初始化
                    if new_head is None:
                        new_head = node
                        point = new_head
                    # 若新链表头结点不为None，则进行尾部连接操作
                    else:
                        point.next = node
                        point = point.next
                    point1 = point1.next
                # 当val2小于val1时，将val2构造成结点放置新链表尾部
                else:
                    node = ListNode(val2)
                    if new_head is None:
                        new_head = node
                        point = new_head
                    else:
                        point.next = node
                        point = point.next
                    point2 = point2.next
            # 当l2已遍历完而l1未遍历完时，直接将point1指针连接至新链表尾部
            elif point1:
                # 当point指针为None时，代表l2链表为空，此时直接将new_head初始化为point1即可
                if point:
                    point.next = point1
                # 若point不为空，则直接进行尾部连接操作
                else:
                    new_head = point1
                # 直接退出当前循环
                break
            # 当l1已遍历完而l2未遍历完时，直接将point2指针连接至新链表尾部
            elif point2:
                # 当point指针为None时，代表l1链表为空，此时直接将new_head初始化为point2即可
                if point:
                    point.next = point2
                # 若point不为空，则直接进行尾部连接操作
                else:
                    new_head = point2
                # 直接退出当前循环
                break
        return new_head


def get_value(head):
    point = head
    result = []
    while point:
        val = point.val
        result.append(val)
        point = point.next
    return result


if __name__ == '__main__':
    # 测试用例
    l1 = ListNode(1)
    l1.next = ListNode(2)
    l1.next.next = ListNode(4)
    l2 = ListNode(1)
    l2.next = ListNode(3)
    l2.next.next = ListNode(4)
    solution = Solution()
    head = solution.mergeTwoLists(l1, l2)
    print(get_value(head))