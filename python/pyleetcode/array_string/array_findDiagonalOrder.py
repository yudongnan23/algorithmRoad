"""

    给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
    举例：
        - 输入：[
                [1, 2, 3, 4],
                [5, ,6 ,7 ,8],
                [9, 10, 11, 12],
                [13, 14, 15, 16]
                ]
        - 输出：[1, 2, 5, 9, 6, 3, 4, 7, 10, 13, 14, 11, 8, 12, 15, 16]
"""


"""

    思路：定义一个规则遍历二维数组所有元素，遍历方向根据横纵下标的改变进行相应更改，当遍历来到二维数组矩阵的边界时需要改变方向。详见代码中的注释

"""


class Solution:
    def findDiagonalOrder(self, matrix):
        # 定义遍历结果的列表，初始化为空列表
        result = []
        if len(matrix) == 0:
            return result
        # 定义纵下标i，横下标j
        i = 0
        j = 0
        # 定义遍历方向，可知遍历方向有上下两个状态，初始化为up上状态
        token = "up"
        while i < len(matrix) and j < len(matrix[0]):
            # 遍历方向为up且横下标达到极值时，即达到二维数组的右上角的元素，下一步需要走向当前元素的下面的元素，即横下标不变，纵下标加1，并且遍历方向改为down，此情况优先级最高
            if j == len(matrix[0]) - 1 and token == "up":
                result.append(matrix[i][j])
                i = i + 1
                token = "down"
            # 遍历方向为上且纵下标为0，即遍历来到二维矩阵的上边时，下一步需要走向其右边的元素，即纵下标不变，横下标加1，并且将遍历方向改为down，此情况优先级次之
            elif i == 0 and token == "up":
                result.append(matrix[i][j])
                j = j + 1
                token = "down"
            # 遍历方向为下且纵下标到达极值，即遍历来到左下角元素时，下一步需要走向当前元素右边的元素，即横下标加1，纵下标加1，并将遍历方向改为up
            elif i == len(matrix) - 1 and token == "down":
                result.append(matrix[i][j])
                j = j + 1
                token = "up"
            # 遍历方向为下且横下标为0，即遍历来到二维矩阵的左边时，下一步需要走向其下面的元素，即横下标不变，纵下标加1，并且将遍历方向更改为up
            elif j == 0 and token == "down":
                result.append(matrix[i][j])
                i = i + 1
                token = "up"
            # 此情况代表遍历在矩阵中间或者在上一步完成改向，只需要按照遍历方向将横纵下标进行相应的改变即可
            else:
                if token == "down":
                    result.append(matrix[i][j])
                    i = i + 1
                    j = j - 1
                elif token == "up":
                    result.append(matrix[i][j])
                    i = i - 1
                    j = j + 1
        return result


if __name__ == '__main__':
    # 测试用例
    nums = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]
    ss = Solution()
    print(ss.findDiagonalOrder(nums))
