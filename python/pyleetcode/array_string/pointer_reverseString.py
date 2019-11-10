'''

    编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
    不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
    你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

'''

class Solution:
    def reverseString(self, s):
        if len(s) == 0:
            pass
        else:
            font_point = 0
            back_point = len(s) - 1
            while back_point > font_point:
                temp = ""
                temp = s[back_point]
                s[back_point] = s[font_point]
                s[font_point] = temp
                font_point += 1
                back_point -= 1
