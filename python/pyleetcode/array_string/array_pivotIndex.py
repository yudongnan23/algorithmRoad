"""

    给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
    我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
    如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

"""


"""

    思路：定义一个指针，该指针初始化为0，再定义指针指向元素的边所有元素的和fe_sum，和指针指向元素右边所有元素之和两个变量be_sum，对数组进行遍历
         ，每遍历一个元素fe_sum加上遍历过的元素，be_sum减去遍历过的元素，直到fe_sum等于be_sum

"""


class Solution:
    def pivotIndex(self, nums):
        if nums == []:
            return -1
        # 数组指针，初始化为0
        pre = 0
        # 指针左边所有元素之和
        fe_sum = 0
        # 指针右边所有元素之和
        be_sum = 0
        while pre < len(nums):
            # 当指针指向第一个元素时，对be_sum进行求值
            if pre == 0:
                for i in range(1, len(nums)):
                    be_sum = be_sum + nums[i]
            # 当be_sum与be_sum相等时，返回指针指向元素的索引
            if fe_sum == be_sum:
                return pre
            else:
                # fe_sum加上遍历过的元素
                fe_sum = fe_sum + nums[pre]
                try:
                    # be_sum减去遍历过的元素，异常捕捉是为了防止指针来到数组最后一个元素还没有找到新索引，指针继续往前会判抛出异常
                    be_sum = be_sum - nums[pre + 1]
                except:
                    pass
            pre += 1
        return -1
