'''

    给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。

'''


class Solution:
    def findDiagonalOrder(self, matrix):
        result = []
        if len(matrix) == 0:
            return result
        i = 0
        j = 0
        token = "up"
        while i < len(matrix) and j < len(matrix[0]):
            if j == len(matrix[0]) - 1 and token == "up":
                result.append(matrix[i][j])
                i = i + 1
                token = "down"
            elif i == 0 and token == "up":
                result.append(matrix[i][j])
                j = j + 1
                token = "down"
            elif i == len(matrix) - 1 and token == "down":
                result.append(matrix[i][j])
                j = j + 1
                token = "up"
            elif j == 0 and token == "down":
                result.append(matrix[i][j])
                i = i + 1
                token = "up"
            else:
                if token == "down":
                    result.append(matrix[i][j])
                    i = i + 1
                    j = j - 1
                elif token == "up":
                    result.append(matrix[i][j])
                    i = i -1
                    j = j + 1
        return result

if __name__ == '__main__':
    nums = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]
    ss = Solution()
    print(ss.findDiagonalOrder(nums))