'''

   实现 strStr() 函数。
    给定一个 haystack 字符串和一个 needle 字符串，
    在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
    如果不存在，则返回  -1

'''

class Solution:
    def strStr(self, haystack: str, needle: str):
        if needle == '':
            return 0
        i = 0
        len_needle = len(needle)
        while len(haystack) - i > len_needle:
            if haystack[i:i+len_needle] == needle:
                return i
            i += 1
        return -1

if __name__ == '__main__':
    ss = "aaaaa"
    s = 'bba'
    object = Solution()
    print(object.strStr(ss, s))