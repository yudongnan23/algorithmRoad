"""

    给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素

"""


"""

    思路：定义四个函数，up函数负责对二维数组矩阵的上边元素进行遍历，right函数负责对二维数组矩阵的右边元素进行遍历，
         down函数负责对二维矩阵的下边元素进行遍历，left函数负责对二维矩阵的左边元素进行比遍历，每个函数遍历完成后将遍历
         过的的元素在二维矩阵中删除，再将二维矩阵传递给下一个函数进行遍历，四个函数遍历顺序为：
            up - right - down - left - up

"""


class Solution:
    def spiral_order(self, matrix):
        result = []
        if len(matrix) == 0:
            return result
        # 调用up函数对二维矩阵进行遍历
        result = self.up(matrix, result)
        return result

    def up(self, nums, result):
        # 当传入的二维矩阵为空时，返回当前遍历结果列表result
        if len(nums) == 0 or len(nums[0]) == 0:
            return result获得二维矩阵的上边所有元素
        #
        result_item = nums[0]
        # 获得需要传入下一个函数的二维矩阵，即删除二维矩阵中所有上边元素
        new_nums = nums[1:]
        # 将二维矩阵的上边的所有元素添加进遍历结果列表
        result = result + result_item
        # 返回下一个函数调用的结果，需要传入当前遍历结果列表
        return self.right(new_nums, result)

    def right(self, nums, result):
        # 当传入的二维矩阵为空时，返回当前遍历结果列表result
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        # 获得二维矩阵的右边所有的元素
        result_item = [item[-1] for item in nums]
        # 获得需要传入下一个函数的二维矩阵，即删除二维矩阵的右边的所有元素
        new_nums = [num[0:len(num)-1] for num in nums]
        # 将二维矩阵的右边的所有元素添加进遍历结果列表
        result = result + result_item
        # 返回下一个函数调用的结果，需要传入当前遍历结果列表
        return self.down(new_nums, result)

    def down(self, nums, result):
        # 当传入的二维矩阵为空时，返回当前遍历结果列表result
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        # 获得二维矩阵的下边所有的元素
        result_item = list(reversed(nums[-1]))
        # 获得需要传入下一个函数的二维矩阵，即删除二维矩阵的下边的所有元素
        new_nums = nums[0:len(nums)-1]
        # 将二维矩阵的下边的所有元素添加进遍历结果列表
        result = result + result_item
        # 返回下一个函数调用的结果，需要传入当前遍历结果列表
        return self.left(new_nums, result)

    def left(self, nums, result):
        # 当传入的二维矩阵为空时，返回当前遍历结果列表result
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        # 获得二维矩阵的左边所有的元素
        result_item = list(reversed([item[0] for item in nums]))
        # 获得需要传入下一个函数的二维矩阵，即删除二维矩阵的左边的所有元素
        new_nums = [num[1:] for num in nums]
        # 将二维矩阵的左边的所有元素添加进遍历结果列表
        result = result + result_item
        # 返回下一个函数调用的结果，需要传入当前遍历结果列表
        return self.up(new_nums, result)


if __name__ == '__main__':
    ss = Solution()
    num = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]
    print(ss.spiral_order(num))