'''

    给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素

'''


class Solution:
    def spiral_order(self, matrix):
        result = []
        if len(matrix) == 0:
            return result
        result = self.up(matrix, result)
        return result

    def up(self, nums, result):
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        result_item = nums[0]
        new_nums = nums[1:]
        result = result + result_item
        return self.right(new_nums, result)

    def right(self, nums, result):
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        result_item = [item[-1] for item in nums]
        new_nums = [num[0:len(num)-1] for num in nums]
        result = result + result_item
        return self.down(new_nums, result)

    def down(self, nums, result):
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        result_item = list(reversed(nums[-1]))
        new_nums = nums[0:len(nums)-1]
        result = result + result_item
        return self.left(new_nums, result)

    def left(self, nums, result):
        if len(nums) == 0 or len(nums[0]) == 0:
            return result
        result_item = list(reversed([item[0] for item in nums]))
        new_nums = [num[1:] for num in nums]
        result = result + result_item
        return self.up(new_nums, result)


if __name__ == '__main__':
    ss = Solution()
    nums = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]
    print(ss.spiral_order(nums))