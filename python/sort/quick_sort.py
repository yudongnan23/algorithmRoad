"""

快速排序

    算法描述
        快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

    算法描述
        快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

            从数列中挑出一个元素，称为 “基准”（pivot）；
            重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
            递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
"""


class Solution:
    def quick_sort(self, nums: list) -> list:
        """ sort the array by the quick sort method

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0:
            return nums

        left = 0
        right = length - 1
        self.sort(nums, left, right)

        return nums

    def sort(self, nums: list, left: int, right: int):
        """

        :param nums:
        :param left:
        :param right:
        :return:
        """
        if left < right:
            partition_index = self.get_partition_index(nums, left, right)

            self.sort(nums, left, partition_index - 1)
            self.sort(nums, partition_index + 1, right)

    def get_partition_index(self, nums: list, left: int, right: int) -> int:
        """

        :param nums:
        :param left:
        :param right:
        :return:
        """
        pivot = left
        cur_index = pivot + 1

        for variant_index in range(left, right + 1):
            if nums[variant_index] < nums[pivot]:
                self.swap(nums, cur_index, variant_index)
                cur_index += 1

        self.swap(nums, pivot, cur_index - 1)

        return cur_index - 1

    def swap(self, nums: list, i: int, j: int) -> None:
        """

        :param nums:
        :param i:
        :param j:
        :return:
        """
        temp = nums[j]
        nums[j] = nums[i]
        nums[i] = temp


if __name__ == '__main__':
    nums = [1, 2, 6, 8, 0, 0, 10, 100, 200]
    solution = Solution()
    print(solution.quick_sort(nums))