'''

    给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

'''

class Solution:
    def rotate(self, nums, k):
        if len(nums) == 0:
            return nums
        new_nums = nums[k+1:] + nums[0:k+1]
        return new_nums


if __name__ == '__main__':
    ss = Solution()
    nums = [1, 2, 3, 4, 5, 6, 7]
    k = 3
    print(ss.rotate(nums, k))