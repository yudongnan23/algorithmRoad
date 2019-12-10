"""

    设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
    如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的

"""
"""

    思路：自定义一个函数用于获取当前链表的长度
        get：获得链表的长度，判断索引是否正确，正确则定义一个指针，让指针走index步，即可得到index位置上结点的值
        addAtHead：根据传入的结点值初始化一个结点，判断头结点是否为空，若为空，将头结点指针指向初始化的结点即可；否则将该结点的next
                   指定为头结点，再将头结点指针指向初始化的结点。
        addAtTail：根据初入的结点初始化一个结点，判断头结点是否是空，若为空，将头结点指针指向初始化的结点即可；否则，定义一个动态指针
                   让指针走到链表的表尾，然后将指针指向的结点的next指向初始化的结点。
        addAtIndex：首先根据传入的结点值初始化一个结点，然后获取当链表的长度，判断传入的index是否小于或等于0，若是，执行AddAtHead操作；
                    若不是，判断index是否小于链表的长度，若是，则定义一个动态指针，初始化为链表的头部，让该指针走index-1步，将初始的结点
                    的next指向当前指针指向的结点的next结点，再将当前指针指向的结点的next指向初始化的结点；如不是，判断index是否等于
                    链表的长度，若是，则进行AddAtTail操作；若不是，则代表index大于链表的长度，则不进行任何操作
        deleteAtIndex： 首先获取链表的长度，index是否有效，若无效，则不进行任何操作；若有效，则初始化一个结点值为-1（自定义）的结点，
                        将初始化的结点连接至表头，定义一个动态指针，指向初始化的结点，这样，在执行删除操作时，指针的next结点可以是链表中的任何结点
                        便于执行删除操作。让指针走index步，此时指针的结点的next结点刚好为第index个结点，将指针指向的结点next指向
                        指针指向的next结点的next结点，这样就达到了删除第index个节点的效果。
"""


class Node:
    def __init__(self, val=None):
        self.val = val
        self.next = None


class MyLinkedList:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.head = Node()

    def get(self, index):
        """
        Get the value of the index-th node in the linked list. If the index is invalid, return -1.
        """
        length = self.get_length()
        # 判断index是否有效，即大于或等于0，小于链表长度
        if index in range(0, length):
            count = 0
            # 定义动态指针，指向头结点
            point = self.head
            # 让指针走index步
            while count < index:
                point = point.next
                count += 1
            return point.val
        else:
            return -1

    def addAtHead(self, val):
        """
        Add a node of value val before the first element of the linked list. After the insertion, the new node will be
        the first node of the linked list.
        """
        # 初始化新的结点
        node = Node(val)
        # 判断头结点是否为空
        if self.head.val is None:
            # 将结点的头指针指向初始化的新结点
            self.head = node
        else:
            # 将初始化的新结点的next指向头结点，再将头结点指针指向初始化的新结点
            temp = self.head
            self.head = node
            self.head.next = temp

    def addAtTail(self, val):
        """
        Append a node of value val to the last element of the linked list.
        """
        # 初始化新的结点
        node = Node(val)
        # 判断头结点是否为空
        if self.head.val is None:
            # 将结点的头指针指向初始化的新结点
            self.head = node
        else:
            # 定义动态指针，初始化为头结点
            point = self.head
            # 让指针走到表尾
            while point.next:
                point = point.next
            # 指针指向的表尾结点的next指向初始化的新结点
            point.next = node

    def addAtIndex(self, index, val):
        """
        Add a node of value val before the index-th node in the linked list. If index equals to the length of linked
        list, the node will be appended to the end of linked list. If index is greater than the length, the node will
        not be inserted.
        """
        node = Node(val)
        # 调用函数获取链表的长度
        length = self.get_length()
        # 判断索引是否在1-length范围之内
        if index in range(1, length):
            count = 0
            point = self.head
            # 让指针走index-1步
            while count < index - 1:
                point = point.next
                count += 1
            # 执行插入操作
            node.next = point.next
            point.next = node
        # 若index小于或等于0，执行addAtHead操作
        elif index <= 0:
            self.addAtHead(val)
        # 若index等于链表的长度，执行addAtTail操作
        elif index == length:
            self.addAtTail(val)
        # 其他情况，即index大于链表长度，不执行任何操作
        else:
            pass

    def deleteAtIndex(self, index):
        """
        Delete the index-th node in the linked list, if the index is valid.
        """
        length = self.get_length()
        pre_node = Node(-1)
        # 将新初始化的结点的next指向头结点
        pre_node.next = self.head
        # 动态指针指向新初始化的结点，保持了该指针可以始终在需要删除的结点的前一个节点
        point = pre_node
        # 判断index是否有效，在0-length范围内，否则不进行任何操作
        if index in range(0, length):
            count = 0
            # 让指针走index步
            while count < index:
                point = point.next
                count += 1
            # 执行删除操作
            point.next = point.next.next
            # 若index等于0，那么删除的结点会是头结点，需要将头结点指针指向动态指针的next结点
            if index == 0:
                self.head = point.next
        else:
            pass

    # 额外添加的方法，用于获取链表的长度
    def get_length(self):
        length = 0
        if self.head.val is None:
            return length
        point = self.head
        while point:
            point = point.next
            length += 1
        return length

    # 额外添加的方法，用于将链表转换为列表
    def get_val_to_list(self):
        result = []
        point = self.head
        while point:
            result.append(point.val)
            point = point.next
        return result


# 测试用例
if __name__ == '__main__':
    linklist = MyLinkedList()
    linklist.addAtIndex(0, 10)
    linklist.addAtIndex(0,20)
    linklist.addAtIndex(1, 30)
    print(linklist.get_val_to_list())
    print(linklist.get(0))
