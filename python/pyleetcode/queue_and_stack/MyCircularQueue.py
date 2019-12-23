"""

    设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
    循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
    你的实现应该支持如下操作：

        MyCircularQueue(k): 构造器，设置队列长度为 k 。
        Front: 从队首获取元素。如果队列为空，返回 -1 。
        Rear: 获取队尾元素。如果队列为空，返回 -1 。
        enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
        deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
        isEmpty(): 检查循环队列是否为空。
        isFull(): 检查循环队列是否已满。

"""


"""

    思路：使用一个列表来实现该循环队列，需要初始化该列表，该列表的长度为k，队列的队首元素在列表中的索引-1，队列的队尾元素在列中的索引-1
            isEmpty()：判断队首元素索引与队尾元素是否都为-1，若是，则返回True，否则返回False
            isFull()：判断队尾元素的索引加1对队列长度求余是否等于队首元素索引，若是，则返回True，否则返回False
            Front：调用实例方法isEmpty()方法判断队列是否为空，若是，则返回-1，否则返回队首元素
            Rear：调用实例方法isEmpty()方法判断队列是否为空，若是，则返回-1，否则返回队尾元素
            enQueue：调用isFull()方法判断队列是否已满，若是，返回False；否则，判断队首元素的索引是否等于-1，若是，代表队列为空，将队首元素索引改为0；使用队尾元素的索引加1对
                     队列的长度进行求余得到队尾元素新的索引，向该索引插入新的元素即可。
            deQueue：调用isEmpty()判断队列是否为空，若为空，返回False；判断队尾元素与队首元素的索引是否相等，若相等，代表队列中只有一个元素，将队首元素与队尾元素都置为-1，
                     若不相等，将队首元素的索引加1对队列长度进行求余即为新队首元素的索引。

"""


class MyCircularQueue:

    def __init__(self, k: int):
        """
        Initialize your data structure here. Set the size of the queue to be k.
        """
        self.__length = k
        self.__queue = [None] * k
        self.__start = -1
        self.__end = -1

    def enQueue(self, value: int) -> bool:
        """
        Insert an element into the circular queue. Return true if the operation is successful.
        """
        enqueue_result = True
        if not self.isFull():
            if self.__start == -1:
                self.__start = 0
            self.__end = (self.__end + 1) % self.__length
            self.__queue[self.__end] = value
        else:
            enqueue_result =  False
        return enqueue_result

    def deQueue(self) -> bool:
        """
        Delete an element from the circular queue. Return true if the operation is successful.
        """
        dequeue_result = True
        if not self.isEmpty():
            if self.__end == self.__start:
                self.__end = -1
                self.__start = -1
            else:
                self.__start = (self.__start + 1) % self.__length
        else:
            dequeue_result = False
        return dequeue_result

    def Front(self) -> int:
        """
        Get the front item from the queue.
        """
        if self.isEmpty():
            return -1
        val = self.__queue[self.__start]
        return val

    def Rear(self) -> int:
        """
        Get the last item from the queue.
        """
        if self.isEmpty():
            return -1
        val = self.__queue[self.__end]
        return val

    def isEmpty(self) -> bool:
        """
        Checks whether the circular queue is empty or not.
        """
        return self.__start == -1 and self.__end == -1

    def isFull(self) -> bool:
        """
        Checks whether the circular queue is full or not.
        """
        return (self.__end + 1) % self.__length == self.__start
