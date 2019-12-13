"""

    给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
    最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
    你可以假设除了整数 0 之外，这个整数不会以零开头。

"""


"""

    思路：从列表的最后一个元素开始遍历，当遇到9时，需要对当前元素进行改1操作，进行下一个元素的遍历；当遇到的不是9时，只需要将当前遍历元素加1，然后结束遍历，返回列表即可
    
"""


class Solution:
    def plusOne(self, digits):
        if len(digits) == 0:
            return []
        # 获得列表元素最大索引
        index = len(digits) - 1
        # 对列表元素进行遍历
        while index >= 0:
            # 当遇到9时
            if digits[index] == 9:
                # 当前元素直接改0
                digits[index] = 0
                # 如果数组元素最大索引为0，即列表只存在一个元素9，将1添加至列表头部，直接返回列表即可
                if index == 0:
                    result = [1]+digits
                    return result
                index = index - 1
            else:
                # 遍历得到的元素不是9时，将当前元素加1直接结束遍历
                digits[index] = digits[index] + 1
                break
        return digits


if __name__ == '__main__':
    # 测试用例
    ss = Solution()
    print(ss.plusOne([9, 9]))
