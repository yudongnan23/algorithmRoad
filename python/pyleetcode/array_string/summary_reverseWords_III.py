'''

    给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序

'''
class Solution:
    def reverseWords(self, s):
        s_list = s.split(" ")
        index = 0
        result_list = []
        s_list_length = len(s_list)
        while index < s_list_length:
            if s_list[index] != "":
                result_list.append(s_list[index][::-1])
            index += 1
        return " ".join(result_list)

    # def reverse(self, string):
    #     result_string = ""
    #     for s in string:
    #         result_string = s + result_string
    #     return result_string

if __name__ == '__main__':
    s =  "Let's take LeetCode contest"
    solution = Solution()
    print(solution.reverseWords(s))
