"""
堆排序

    堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

    算法描述
        将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
        将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
        由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成

"""


class Solution:
    def heap_sort(self, nums: list) -> list:
        """

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0 or length == 1:
            return nums

        self.build_heap(nums, length)

        variant_index = length - 1
        while variant_index > 0:
            self.swap(nums, 0, variant_index)
            self.heapify(nums, variant_index, 0)

            variant_index -= 1

        return nums

    def build_heap(self, nums: list, length: int):
        """

        :param nums:
        :param length:
        :return:
        """
        last_node = length - 1
        parent_node = (last_node - 1) // 2

        index = parent_node

        while index >= 0:
            self.heapify(nums, length, index)
            index -= 1

    def heapify(self, nums: list, length: int, cur_index: int):
        """

        :param nums:
        :param length:
        :param cur_index:
        :return:
        """
        if cur_index >= length:
            return

        s1 = 2 * cur_index + 1
        s2 = 2 * cur_index + 2
        max_index = cur_index

        if s1 < length and nums[s1] > nums[max_index]:
            max_index = s1

        if s2 < length and nums[s2] > nums[max_index]:
            max_index = s2

        if max_index != cur_index:
            self.swap(nums, max_index, cur_index)
            self.heapify(nums, length, max_index)

    def swap(self, nums: list, index1: int, index2: int):
        """

        :param nums:
        :param index1:
        :param index2:
        :return:
        """
        temp = nums[index1]
        nums[index1] = nums[index2]
        nums[index2] = temp


if __name__ == '__main__':
    array = [1, 0, 2, 8, 0, 100, 2, 0, 100]
    solution = Solution()
    print(solution.heap_sort(array))