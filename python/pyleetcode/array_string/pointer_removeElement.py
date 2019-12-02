'''

    给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

    不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

    元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

'''

class Solution:
    def removeElement(self, num, val):
        result = 0
        if len(num) == 0:
            return result
        font_point = 0
        bcak_point = len(num) - 1
        count = 0
        while font_point < bcak_point:
            if num[font_point] == val:
                temp = num[bcak_point]
                num[bcak_point] = num[font_point]
                num[font_point] = temp
                bcak_point -= 1
            else:
                count += 1
                font_point += 1
        if num[font_point] != val:
            count += 1
        result = len(num[0:count])
        return result


if __name__ == '__main__':
    num = [0, 1, 2, 2, 3, 0, 4, 2]
    val = 2
    ss = Solution()
    print(ss.removeElement(num, val))
