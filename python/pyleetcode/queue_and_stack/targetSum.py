"""

    给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
    返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

"""


"""

    思路： 采用深度优先搜索算法

"""


class Solution1:
    def findTargetSumWays(self, nums, S):
        """ find the target sum by nums, and get the total ways to get that

        :param nums:
            :type: list[int]
        :param S:
            :type: int
        :return:
            :rtype: int
        """
        ways = 0
        if len(nums) == 0:
            return ways

        multiple = [1, -1]
        stack = [[nums[0], 0], [-nums[0], 0]]
        max_num_index = len(nums) - 1
        while stack:
            cur_sum, cur_num_index = stack.pop()

            if cur_num_index == max_num_index:
                if cur_sum == S:
                    ways += 1
                continue

            next_num_index = cur_num_index + 1
            for value in multiple:
                next_sum = cur_sum + nums[next_num_index] * value
                stack_new_item = (next_sum, next_num_index)

                if next_num_index == max_num_index:
                    if next_sum == S:
                        ways += 1
                    continue

                stack.append(stack_new_item)

        return ways


class Solution:
    def findTargetSumWays(self, nums, S):
        """ find the target sum by nums, and get the total ways to get that

        :param nums:
            :type: list[int]
        :param S:
            :type: int
        :return:
            :rtype: int
        """
        cur_sum = 0
        ways = 0
        if len(nums) == 0:
            return ways

        ways = self.dfs(nums, cur_sum, S)

        return ways

    def dfs(self, nums, cur_sum, S):
        """ depth first search the target num

        :param nums:
            :type: list[int]
        :param cur_sum:
            :type: int
        :param S:
            :type: int
        :return:
            :rtype: int
        """
        ways = 0
        nums_length = len(nums)
        if nums_length == 0:
            if cur_sum == S:
                return 1
            return 0

        multiple = [1, -1]
        for value in multiple:
            new_num = nums[nums_length - 1]
            new_nums = nums[0:nums_length - 1]
            new_sum = cur_sum + new_num * value
            ways = ways + self.dfs(new_nums, new_sum, S)

        return ways


if __name__ == '__main__':
    nums = [1, 1, 1, 1, 1]
    S = 3
    solution = Solution()
    print(solution.findTargetSumWays(nums, S))
