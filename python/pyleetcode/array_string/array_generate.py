'''

    给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

'''

class Solution:
    def generate(self, numRows):
        result = []
        if numRows == 0:
            return result
        for i in range(numRows):
            if i == 0:
                result = result + [[1]]
            elif i == 1:
                result = result + [[1, 1]]
            else:
                cur_nums = result[-1]
                new_nums = [1]
                for j in range(len(cur_nums)-1):
                    new_nums.append(cur_nums[j] + cur_nums[j+1])
                new_nums.append(1)
                result.append(new_nums)
        return result
if __name__ == '__main__':
    ss = Solution()
    num = 3
    print(ss.generate(num))
