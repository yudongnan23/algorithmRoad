"""

桶排序
    桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
    桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

    算法描述11
        设置一个定量的数组当作空桶；
        遍历输入数据，并且把数据一个一个放到对应的桶里去；
        对每个不是空的桶进行排序；
        从不是空的桶里把排好序的数据拼接起来。
"""


class Solution:
    def bucket_sort(self, nums: list) -> list:
        """ sort the array by the bucket sort method

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0:
            return nums

        max_value, min_value = self.get_max_and_min_value(nums)

        bucket_size = 5




    def get_max_and_min_value(self, nums: list) -> (int, int):
        """ get the max value and min value in the nums

        :param nums:
        :return:
        """
        max_value = 0
        min_value = 0

        for value in nums:
            if value < min_value:
                min_value = value
                continue

            if value > max_value:
                max_value = value
                continue

        return max_value, min_value
