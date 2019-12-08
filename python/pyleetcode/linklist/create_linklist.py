"""

    设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
    如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的

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
        if index in range(0, length):
            count = 0
            point = self.head
            while count < index:
                point = point.next
                count += 1
            return point.val
        else:
            return -1

    def addAtHead(self, val):
        """
        Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
        """
        node = Node(val)
        if self.head.val is None:
            self.head = node
        else:
            temp = self.head
            self.head = node
            self.head.next = temp

    def addAtTail(self, val):
        """
        Append a node of value val to the last element of the linked list.
        """
        node = Node(val)
        if self.head.val is None:
            self.head = node
        else:
            point = self.head
            while point.next:
                point = point.next
            point.next = node


    def addAtIndex(self, index, val):
        """
        Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
        """
        node = Node(val)
        length = self.get_length()
        if index in range(1, length):
            count = 0
            point = self.head
            while count < index - 1:
                point = point.next
                count += 1
            node.next = point.next
            point.next = node
        elif index <= 0:
            self.addAtHead(val)
        elif index == length:
            self.addAtTail(val)
        else:
            pass

    def deleteAtIndex(self, index):
        """
        Delete the index-th node in the linked list, if the index is valid.
        """
        length = self.get_length()
        pre_node = Node(-1)
        pre_node.next = self.head
        point = pre_node
        if index in range(0, length):
            count = 0
            while count < index:
                point = point.next
                count += 1
            point.next = point.next.next
            if index == 0:
                self.head = point.next
        else:
            pass

    def get_length(self):
        length = 0
        if self.head.val is None:
            return length
        point = self.head
        while point:
            point = point.next
            length += 1
        return length

    def get_val_to_list(self):
        result = []
        point = self.head
        while point:
            result.append(point.val)
            point = point.next
        return result

if __name__ == '__main__':
    linklist = MyLinkedList()
    linklist.addAtIndex(0, 10)
    linklist.addAtIndex(0,20)
    linklist.addAtIndex(1, 30)
    print(linklist.get_val_to_list())
    print(linklist.get(0))
