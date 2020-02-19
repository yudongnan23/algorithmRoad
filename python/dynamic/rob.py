"""

    你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
    如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额

"""


class Solution:
    def rob(self, nums: list) -> int:
        """

        :param nums:
        :return:
        """
        length = len(nums)
        if length == 0:
            return 0
        if length == 1:
            return nums[0]
        if length == 2:
            return max(nums)
        dp = [0] * length
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])
        max_value = 0
        for i in range(2, length):
            dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])

            max_value = max(dp[i], max_value)

        return max_value


if __name__ == '__main__':
    array = [1, 10, 9, 1, 3]
    solution = Solution()
    print(solution.rob(array))
