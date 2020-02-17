"""

    一个机器人位于一个 m x n 网格的左上角）。机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。问总共有多少条不同的路径？

"""

class Solution:
    def total_path(self, m: int, n: int)->int:
        f = [[0]*n] * m
        i, j = 0, 0
        for i in range(m):
            for j in range(n):
                if i == 0 or j == 0:
                    f[i][j] = 1
                else:
                    f[i][j] = f[i - 1][j] + f[i][j - 1]
        return f[m - 1][n - 1]

if __name__ == '__main__':
    solution = Solution()
    print(solution.total_path(2, 2))