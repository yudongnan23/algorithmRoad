"""

插入排序
    插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。

    算法描述
        从第一个元素开始，该元素可以认为已经被排序；
        取出下一个元素，在已经排序的元素序列中从后向前扫描；
        如果该元素（已排序）大于新元素，将该元素移到下一位置；
        重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
        将新元素插入到该位置后；
        重复步骤2~5。

    算法分析
        插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。

"""


class Solution:
    def inser_sort(self, nums: list) -> list:
        """ sort the array by inserting method

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0:
            return nums

        for cur_index in range(length):
            variant_index = cur_index
            temp = nums[variant_index]

            while variant_index > 0:
                if temp < nums[variant_index - 1]:
                    nums[variant_index] = nums[variant_index - 1]
                    variant_index = variant_index - 1
                else:
                    break

            nums[variant_index] = temp

        return nums


if __name__ == '__main__':
    array = [0, 1, 2, 32, 0, 1, 3, 8]
    solution = Solution()
    print(solution.inser_sort(array))