class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def make_link_list(number: int) -> ListNode:
    """
    :param number:
    :return:
    """
    print(number)
    l = ListNode(number % 10, None)
    head = l

    length = len(str(number))
    i = 2
    while i <= length:
        val = number % pow(10, i)
        if len(str(val)) < len(str(pow(10, i))) - 1:
            val = 0
        l.next = ListNode(int(str(val)[0]), None)
        l = l.next
        i += 1
    return head


def get_number(l: ListNode) -> int:
    """
    :param l:
    :return:
    """
    if l == None:
        return 0
    s = ""
    while l:
        s = str(l.val) + s
        l = l.next
    return int(s)


def make_link_list_by_list(l) -> ListNode:
    """
    :param l:
    :return:
    """
    if len(l) == 0:
        return None

    point = ListNode(l[0])
    head = point
    i = 1
    while i < len(l):
        point.next = ListNode(l[i], None)
        point = point.next
        i += 1
    return head


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        """
        :param l1:
        :param l2:
        :return:
        """
        return make_link_list(get_number(l1) + get_number(l2))


if __name__ == '__main__':
    ll1 = make_link_list_by_list([1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1])
    ll2 = make_link_list_by_list([5, 6, 4])
    s = Solution()
    print(get_number(s.addTwoNumbers(ll1, ll2)))