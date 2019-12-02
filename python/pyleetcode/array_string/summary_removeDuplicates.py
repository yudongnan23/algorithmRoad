"""

    给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
    不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
    不需要考虑超出新长度后面的元素

"""

class Solution:
    def removeDuplicates(self, nums):
        if len(nums) == 0:
            return len(nums)
        left_point = 0
        right_point = 0
        while right_point < len(nums) and left_point < len(nums)-1:
            if nums[right_point] != nums[left_point]:
                temp = nums[right_point]
                nums[right_point] = nums[left_point+1]
                nums[left_point+1] = temp
                left_point += 1
            right_point += 1
        return left_point+1


if __name__ == '__main__':
    solution = Solution()
    nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
    print(solution.removeDuplicates(nums))
