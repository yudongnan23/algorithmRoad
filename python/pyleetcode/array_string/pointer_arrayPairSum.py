'''

    给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，
    使得从1 到 n 的 min(ai, bi) 总和最大。

'''

class Solution:
    def arrayPairSum(self, nums):
        result = 0
        if len(nums) == 0:
            return result
        new_nums = sorted(nums)
        i = 0
        while i <= len(new_nums) - 2:
            result = result + new_nums[i]
            i += 2
        return result