"""

    给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

"""


class Solution:
    def max_sub_array(self, nums: list) -> int:
        """

        :param nums:
        :return:
        """
        length = len(nums)
        if length == 0:
            return 0
        # 开一个大小为length的动态数组，且其元素都初始化为0
        dp = [0] * length
        # 由题意可知，动态数组的第一个元素一定为nums数组的第一个元素
        dp[0] = nums[0]
        # 定义一个最大子序列和，初始化为nums数组的第一个数
        max_sum = nums[0]
        # 从nums数组的第二个元素开始计算到每个元素的最大子序列和
        for i in range(1, length):
            add = dp[i - 1] + nums[i]
            not_add = nums[i]
            dp[i] = max(add, not_add)
            max_sum = max(dp[i], max_sum)
        return max_sum


if __name__ == '__main__':
    array = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    solution = Solution()
    print(solution.max_sub_array(array))
