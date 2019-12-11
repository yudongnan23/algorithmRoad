"""
    在一个给定的数组nums中，总是存在一个最大元素 。
    查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
    如果是，则返回最大元素的索引，否则返回-1。
"""


"""

    思路：找出数组中的最大数，再找出剩余元素的最大数，判断数组最大数是否是剩余元素的最大数的两倍即可。

"""


class Solution:
    def dominantIndex(self, nums) -> int:
        # 若数组为空，直接返回-1即可
        if len(nums) == 0:
            return -1
        # 初始化最大数变量
        max_num = 0
        # 初始化最大数索引变量
        max_index = 0
        # 初始化除最大数元素之外的最大数变量
        lt_max_num = 0
        # 遍历得到数组最大数
        for i in range(len(nums)):
            if nums[i] > max_num:
                max_num = nums[i]
                max_index = i
        try:
            # 求得剩余元素的最大数，即第二大元素
            lt_max_num = max(nums[0:max_index]+nums[max_index+1:])
        except:
            pass
        # 若第二大元素与最大元素同时为0，返回-1，否则返回最大数索引
        if lt_max_num == 0:
            if max_num == 0:
                return -1
            return max_index
        # 若第二大元素等于最大元素，返回-1
        elif max_num == lt_max_num:
            return -1
        # 若第二大元素大于最大数的一半，返回-1
        elif max_num / lt_max_num < 2:
            return -1
        else:
            return max_index
