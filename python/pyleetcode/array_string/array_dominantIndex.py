'''
    在一个给定的数组nums中，总是存在一个最大元素 。

    查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

    如果是，则返回最大元素的索引，否则返回-1。
'''

class Solution:
    def dominantIndex(self, nums) -> int:
        if nums == []:
            return -1
        max_num = 0
        max_index = 0
        lt_max_num = 0
        for i in range(len(nums)):
            if nums[i] > max_num:
                max_num = nums[i]
                max_index = i
        try:
            lt_max_num = max(nums[0:max_index]+nums[max_index+1:])
        except:
            pass
        if lt_max_num == 0:
            if max_num == 0:
                return -1
            return max_index
        elif max_num == lt_max_num:
            return -1
        elif max_num / lt_max_num < 2:
            return -1
        else:
            return max_index