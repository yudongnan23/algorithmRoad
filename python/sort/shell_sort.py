"""

希尔排序
    1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。

    算法描述
        先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：

            选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
            按增量序列个数k，对序列进行k 趟排序；
            每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

"""


class Solution:
    def shell_sort(self, nums: list):
        """ sort the array by shell sort method

        :param nums:
        :return:
        """
        length = len(nums)

        if length == 0:
            return nums

        # 初始化间隔序列第一个值
        gap = length // 2

        # 根据间隔序值列进行间隔插入排序， 直到间隔序列值为0
        while gap > 0:
            # 原始插入排序
            for cur_index in range(gap, length):
                variant_index = cur_index
                temp = nums[variant_index]

                while variant_index > 0:
                    if temp < nums[variant_index - gap]:
                        nums[variant_index] = nums[variant_index - gap]
                        variant_index = variant_index - gap
                    else:
                        break

                nums[variant_index] = temp

            # 获得下一个间隔序列值
            gap = gap // 2

        return nums


if __name__ == '__main__':
    array = [10, 19, 0, 6, 0, 8]
    solution = Solution()
    print(solution.shell_sort(array))