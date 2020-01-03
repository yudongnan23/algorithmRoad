"""

    你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
    每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
    锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
    列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
    字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1

"""


"""

    思路： 典型的广度优先搜索，在一棵n叉树中广度搜索目标数字，搜索从起点"0000"出发，将其一步之内可能出现的"1000","0100","0010","0001"
            "9000","0900","0090","0009"数字加入队列，然后从队列中取出队首元素，将该元素一步之内可能出现的数字再次加入队列，以此类推，
            直到遇到目标数字；需要注意的是，为了避免数字的重复搜索，需要建立一个set集合来记录已经搜索过的数字，如果已搜索，则无需进行任何操作，
            若未搜索，则将该数字同时加入队列与已搜索集合。

"""


from collections import deque


class Solution:
    def openLock(self, deadends, target):
        """ get the step to the target string

        :type deadends: list[str]
        :type target: str
        :rtype int
        """
        # 初始化一个集合，里面记录遍历过的数字
        visit = set(deadends)
        # 创建一个双端队列
        queue = deque()
        start = "0000"

        # 判断起点"0000"是否包含在死亡数组里面，判断目标数字是否包含在死亡数字里面；若是，直接返回-1
        if target in deadends or start in deadends:
            return -1

        # 初始化双端队列，第一个值记录起点，第二个值记录走到该数字需要用到的步数
        queue.append([start, 0])
        # 将起点添加进入已遍历集合
        visit.add(start)

        while queue:
            # 从队列中弹出队首元素，得到一个数字以及走到该数字需要用到的步数
            cur, step = queue.popleft()
            step += 1

            for i in range(len(cur)):
                # 当数字串的其中之一需要进行加1操作时，判断该数字是否是9，若是9，则需要进行改零操作，否则直接在原数字上加1即可
                if int(cur[i]) < 9:
                    up = cur[0:i] + str(int(cur[i]) + 1) + cur[i+1:]
                else:
                    up = cur[0:i] + "0" + cur[i+1:]

                # 当数字串的其中之一需要进行减1操作时，判断该数字是否是0，若是0，则需要进行改9操作，否是直接在原数字上减1即可
                if int(cur[i]) > 0:
                    down = cur[0:i] + str(int(cur[i]) - 1) + cur[i+1:]
                else:
                    down = cur[0:i] + "9" + cur[i+1:]

                # 判断数字串某一数字加1操作后的数字是否存在于已遍历集合，若存在，不进行任何操作，否则将该数字加入队列与已遍历集合
                if up not in visit:
                    if up == target:
                        return step
                    queue.append([up, step])
                    visit.add(up)

                # 判断数字串某一数字减1操作后的数字是否存在于已遍历集合，若存在，不进行任何操作，否则将该数字加入队列与已遍历集合
                if down not in visit:
                    if down == target:
                        return step
                    queue.append([down, step])
                    visit.add(down)
        return -1


if __name__ == '__main__':
    # 测试用例
    deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"]
    target = "8888"
    solution = Solution()
    print(solution.openLock(deadends, target))