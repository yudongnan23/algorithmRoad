""""

    给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围

"""


"""

    思路： 本题采用广度优先的思路解题，遍历二维数组，当遇到"1"时，将该元素改"0"，再将其两个索引放入队列，然后以队列不为空为循环条件，弹出队列中的队首元素索引，对其周围
            元素进行判断，当遇到"1"，改"0"，再次索引入队，直到队列为空，代表一个岛屿结束。直到下一个岛屿的出现执行同样的操作

"""


class Solution:
    def numIslands(self, grid: [[int]]) -> int:
        """ get islands from arrays

        :param grid:
            :type: [[str]]
            example: [
                        ["1", "1", "0", "0", "0"],
                        ["1", "1", "0", "0", "0"],
                        ["0", "0", "1", "0", "0"],
                        ["0", "0", "0", "1", "1"]
                    ]
        :return:
            :rtype: int
        """
        # 当二维数组为空时，及时返回0
        if len(grid) == 0 or len(grid) == 0 and len(grid[0]) == 0:
            return 0
        across_length = len(grid[0])
        vertical_length = len(grid)
        island_nums = 0
        # 对二维数组进行遍历
        for i in range(vertical_length):
            for j in range(across_length):
                # 当遇到"1"时，将该元素改"0"后再对其周围元素尽心处理
                if grid[i][j] == "1":
                    grid[i][j] = "0"
                    island_nums += 1
                    self.__get_island(i, j, vertical_length, across_length, grid)

        return island_nums

    def __get_island(self, vertical: int, across: int, vertical_length: int, across_length: int, grid: [[str]]) -> None:
        """ traversal the array

        :param vertical: current vertical coordinates in the array
        :param across: current abscissa in the array
        :param vertical_length: max vertical coordinates in the array
        :param across_length: max abscissa in the array
        :param grid: array
        :return: None
        """
        # 队列初始化入参的两个索引
        queue = [[vertical, across]]
        # 当队列不为空时，执行该循环
        while queue and queue[0]:
            # 获取队列中的两个索引
            m, n = queue.pop()
            # 判断当前索引指向元素在矩阵中的上面的元素
            if m > 0 and grid[m-1][n] == "1":
                grid[m-1][n] = "0"
                queue.append([m-1, n])
            # 判断当前索引指向元素在矩阵中的下面元素
            if m < vertical_length - 1 and grid[m+1][n] == "1":
                grid[m+1][n] = "0"
                queue.append([m+1, n])
            # 判断当前元素指向元素在矩阵中的左边元素
            if n > 0 and grid[m][n-1] == "1":
                grid[m][n-1] = "0"
                queue.append([m, n-1])
            # 判断当前索引指向元素在矩阵中的右边元素
            if n < across_length - 1 and grid[m][n+1] == "1":
                grid[m][n+1] = "0"
                queue.append([m, n+1])


if __name__ == '__main__':
    solution = Solution()

    # 测试用例
    nums = [["1", "1", "0", "0", "0"], ["1", "1", "0", "0", "0"], ["0", "0", "1", "0", "0"], ["0", "0", "0", "1", "1"]]

    print(solution.numIslands(nums))
