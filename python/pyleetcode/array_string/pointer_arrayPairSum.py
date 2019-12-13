"""

    给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，
    使得从1 到 n 的 min(ai, bi) 总和最大。

"""


"""

    思路：首先需要对数组进行从小到大进行排序，再依次构造两两数对，才能使得所有数对的最小值之和最大

"""


class Solution:
    def arrayPairSum(self, nums):
        result = 0
        if len(nums) == 0:
            return result
        # 对数组进行从小到大排序
        new_nums = sorted(nums)
        i = 0
        # 依次加上两两数对中的最小值
        while i <= len(new_nums) - 2:
            result = result + new_nums[i]
            i += 2
        return result
