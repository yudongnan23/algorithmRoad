'''

    给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0

'''


class Solution:
    def minSubArrayLen(self, s, nums):
        if len(nums) == 0:
            return 0
        left = 0
        right = 0
        sum_all = 0
        nums_length = len(nums)
        min_length = nums_length + 1
        while left < right:
            if right < nums_length and  sum_all < s:
                sum_all = sum_all + nums[right]
                right += 1
            else:
                sum_all = sum_all - nums[left]
                left += 1
            if sum_all >= s and min_length > right - left:
                min_length = right - left
        if min_length == nums_length + 1:
            return 0
        return min_length
