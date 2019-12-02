"""

    给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

"""

class Solution:
    def moveZeroes(self, nums):
        if len(nums) == 0:
            pass
        else:
            left_point = -1
            right_point = 0
            while left_point < len(nums) - 1 and right_point < len(nums):
                if nums[left_point+1] == 0:
                    while right_point < len(nums):
                        if nums[right_point] != 0:
                            temp = nums[right_point]
                            nums[right_point] = nums[left_point+1]
                            nums[left_point+1] = temp
                            left_point += 1
                        right_point += 1
                else:
                    left_point += 1
                    right_point += 1
        print(nums)


if __name__ == '__main__':
    solution = Solution()
    nums = [0, 1, 0, 3, 12]
    solution.moveZeroes(nums)