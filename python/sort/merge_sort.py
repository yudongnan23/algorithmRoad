"""

归并排序
    归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
    将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

    算法描述
        把长度为n的输入序列分成两个长度为n/2的子序列；
        对这两个子序列分别采用归并排序；
        将两个排序好的子序列合并成一个最终的排序序列。

"""


class Solution:
    def merge_sort(self, nums: list) -> list:
        """ sort the array by merge sort method

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0 or length == 1:
            return nums

        middle = length // 2

        return self.merge(self.merge_sort(nums[0: middle]), self.merge_sort(nums[middle: length]))

    def merge(self, left_nums: list, right_nums: list) -> list:
        """

        :param left_nums:
        :param rigth_nums:
        :return:
        """
        new_nums = []

        left_nums_length = len(left_nums)
        right_nums_length = len(right_nums)

        left_nums_index = 0
        right_nums_index = 0

        while left_nums_index < left_nums_length or right_nums_index < right_nums_length:
            if left_nums_index < left_nums_length and right_nums_index < right_nums_length:
                if left_nums[left_nums_index] < right_nums[right_nums_index]:
                    new_nums.append(left_nums[left_nums_index])
                    left_nums_index += 1
                else:
                    new_nums.append(right_nums[right_nums_index])
                    right_nums_index += 1
            elif left_nums_index < left_nums_length:
                new_nums = new_nums + left_nums[left_nums_index: left_nums_length]
                break
            elif right_nums_index < right_nums_length:
                new_nums = new_nums + right_nums[right_nums_index: right_nums_length]
                break

        return new_nums


if __name__ == '__main__':
    array = [0, 10, 0, 100, 1, 8, 9, 20, 30]
    nums1 = [1, 3, 4, 6]
    nums2 = [2, 3, 6, 7, 8]
    solution = Solution()
    print(solution.merge_sort([0, 0, 0, 1, 9, 10, 18, 0]))