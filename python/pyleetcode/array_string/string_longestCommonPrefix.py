'''

    编写一个函数来查找字符串数组中的最长公共前缀。
    如果不存在公共前缀，返回空字符串 ""。

'''
class Solution:
    def longestCommonPrefix(self, strs):
        if len(strs) == 0:
            return ""
        i = 0
        min_length_string = self.min_length(strs)
        while i < min_length_string:
            j = 0
            while j < len(strs):
                if strs[j][i] != strs[0][i]:
                    return strs[j][0:i]
                j += 1
            i += 1
        return strs[0][0:min_length_string]

    def min_length(self, strs):
        if len(strs) == 0:
            return 0
        result = len(strs[0])
        for str in strs:
            if len(str) < result:
                result = len(str)
        return result

if __name__ == '__main__':
    ss = Solution()
    strs = ["flower","flow","flight"]
    print(min(strs))
    print(ss.longestCommonPrefix(strs))

