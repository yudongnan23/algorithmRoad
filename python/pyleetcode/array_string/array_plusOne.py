'''
    给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

    最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

    你可以假设除了整数 0 之外，这个整数不会以零开头。

'''
class Solution:
    def plusOne(self, digits):
        if digits == []:
            return []
        index = len(digits) - 1
        while index >= 0:
            if digits[index] == 9:
                digits[index] = 0
                if index == 0:
                    result = [1]+digits
                    return result
                index = index -1
            else:
                digits[index] = digits[index] + 1
                break
        return digits


if __name__ == '__main__':
    ss = Solution()
    print(ss.plusOne([9, 9]))
