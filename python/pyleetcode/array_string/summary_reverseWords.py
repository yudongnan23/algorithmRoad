'''

    给定一个字符串，逐个翻转字符串中的每个单词

'''
class Solution:
    def reverseWords(self, s):
        s_list = s.split(" ")
        result_list = []
        index = len(s_list) - 1
        while index >= 0:
            if s_list[index] != "":
                result_list.append(s_list[index])
            index -= 1
        return " ".join(result_list)