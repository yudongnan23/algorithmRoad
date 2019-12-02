'''

    给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

    函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

    说明:

        返回的下标值（index1 和 index2）不是从零开始的。
        你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

'''

class Solution:
    def twoSum(self, numbers, target):
        if len(numbers) == 0:
            return None
        font_point = 0
        back_point = len(numbers) - 1
        while font_point < back_point:
            if numbers[font_point] + numbers[back_point] == target:
                return [font_point+1, back_point+1]
            if numbers[font_point] + numbers[back_point] < target:
                font_point += 1
                continue
            elif numbers[font_point] + numbers[back_point] > target:
                back_point -= 1
        return None


if __name__ == '__main__':
    nums = [-1, 0]
    target = -1
    ss = Solution()
    print(ss.twoSum(nums, target))