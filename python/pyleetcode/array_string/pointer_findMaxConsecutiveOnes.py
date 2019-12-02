'''

    给定一个二进制数组， 计算其中最大连续1的个数。

'''


class Solution:
    def findMaxConsecutiveOnes(self, nums):
        if len(nums) == 0:
            return 0
        result = 0
        variant_point = 0
        while variant_point < len(nums):
            while variant_point < len(nums) and nums[variant_point] != 1:
                variant_point += 1
            count = 0
            while variant_point < len(nums) and nums[variant_point] == 1:
                count += 1
                variant_point += 1
            if result < count or result == 0:
                result = count
        return result

if __name__ == '__main__':
    nums = [1, 1, 0, 0, 1, 1, 1]
    ss = Solution()
    print(ss.findMaxConsecutiveOnes(nums))


