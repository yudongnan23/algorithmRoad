""""

    给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
    请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

"""


"""

    思路：定义四个指针，动态指针point，初始化为head，用于遍历整个链表；偶数区头结点指针odd_head，初始化为head.next, 用于指向偶数去头结点；
         偶数区尾结点指针odd_tail，初始化为head.next, 用于指向偶数区的尾结点；奇数区尾结点even_tail,初始化为head结点。
         动态指针遍历链表，当动态指针走奇数步或者走偶数步时，作相应的插入及赋值即可。详见注释         

"""


class Solution:
    def oddEvenList(self, head):
        if head is None:
            return None
        # 指针走的步数，初始化为1
        count = 1
        # 动态指针，用于遍历整个链表
        point = head
        # 偶数区头结点指针，初始化为head.next
        odd_head = head.next
        # 偶数区尾结点指针，初始化为head.next
        odd_tail = head.next
        # 奇数区尾结点指针，初始化为head
        even_tail = head
        while point:
            # 当指针走1步或者两步时，奇偶区已经分明，不需要任何操作，指针往下走即可
            if count == 1 or count == 2:
                point = point.next
                count += 1
                continue
            # 当指针走的步数为偶数时，将指向偶数区的尾结点的指针指向next结点即可，然后指针继续往下走
            if count % 2 == 0:
                odd_tail = odd_tail.next
                point = point.next
                count += 1
                continue
            # 当指针走的步数为奇数时，需要做相应的插入操作
            if count % 2 == 1:
                # 获取动态指针指向的结点，需要将该结点插入奇数区危结点与偶数区头结点之间
                cur = point
                # 动态指针往下走一步
                point = point.next
                # 偶数区尾结点的next结点需要指向动态指针指向的结点
                odd_tail.next = point
                count += 1
                # 需要插入的节点的next指向偶数区头结点
                cur.next = odd_head
                # 奇数区尾结点指向需要插入的结点，至此结点完成插入
                even_tail.next = cur
                # 新插入一个结点，其实是插入在奇数区尾部，需要重新定义奇数区尾结点
                even_tail = cur
        return head
