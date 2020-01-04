"""

    给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

"""


"""

    思路： 深度优先搜索"1"

"""


from collections import deque


class Solution:
    def numIslands(self, gird: [[str]]) -> int:
        """ get the islands nums in the two-dimensional matrix

        :param gird:
        :return:
        """
        nums_island = 0

        if len(gird) == 0 or len(gird[0]) == 0:
            return nums_island

        across_length = len(gird)
        vertical_length = len(gird[0])

        for across in range(across_length):
            for vertical in range(vertical_length):
                if gird[across][vertical] == "1":
                    nums_island += 1
                    gird[across][vertical] = "0"
                    self.__dfs(across, vertical, across_length, vertical_length, gird)

        return nums_island

    def __dfs(self, across: int, vertical: int, across_length: int, vertical_length: int, grid: [[str]]) -> None:
        """ depth first search the "1" which in the two-dimensional matrix, and then change the "1" to the "0"

        :param across:
        :param vertical:
        :param across_length:
        :param vertical_length:
        :param grid:
        :return:
        """
        stack = deque()
        stack.append([across, vertical])
        while stack:
            across_index, vertical_index = stack.popleft()

            if across_index > 0 and grid[across_index - 1][vertical_index] == "1":
                grid[across_index - 1][vertical_index] = "0"
                stack.append([across_index - 1, vertical_index])

            if across_index < across_length - 1 and grid[across_index + 1][vertical_index] == "1":
                grid[across_index + 1][vertical_index] = "0"
                stack.append([across_index + 1, vertical_index])

            if vertical_index > 0 and grid[across_index][vertical_index - 1] == "1":
                grid[across_index][vertical_index - 1] = "0"
                stack.append([across_index, vertical_index - 1])

            if vertical_index < vertical_length - 1 and grid[across_index][vertical_index + 1] == "1":
                grid[across_index][vertical_index + 1] = "0"
                stack.append([across_index, vertical_index + 1])
