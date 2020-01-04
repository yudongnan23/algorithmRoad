"""

    设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

        push(x) -- 将元素 x 推入栈中。
        pop() -- 删除栈顶的元素。
        top() -- 获取栈顶元素。
        getMin() -- 检索栈中的最小元素。

"""


class MinStack:
    def __init__(self):
        """ initialize the stack

        """
        self.__stack = []
        self.__top = None
        self.__min = None

    def push(self, x: int) -> None:
        self.__stack.append(x)
        self.__top = x
        if self.__min is None or x < self.__min:
            self.__min = x

    def pop(self) -> None:
        if len(self.__stack) != 0:
            pop = self.__stack.pop()
            if len(self.__stack) != 0:
                self.__top = self.__stack[-1]
                if pop == self.__min:
                    self.__min = min(self.__stack)
            else:
                self.__top = None
                self.__min = None

    def top(self) -> int:
        return self.__top

    def getMin(self) -> int:
        return self.__min


if __name__ == '__main__':
    stack = MinStack()
    stack.push(-2)
    stack.push(0)
    stack.push(-3)
    print(stack.getMin())
    stack.pop()
    print(stack.top())
    print(stack.getMin())
