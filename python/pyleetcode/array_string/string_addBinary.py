'''

    给定两个二进制字符串，返回他们的和（用二进制表示）。
    输入为非空字符串且只包含数字 1 和 0。
'''


from math import pow


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        if a == '':
            return b
        if b == '':
            return a
        a = a[::-1]
        b = b[::-1]
        new_a = 0
        new_b = 0
        i = 0
        j = 0
        while i < len(a):    # 将字符串a转换为十进制数（int(a, 2)可以将字符串快速转换为十进制数）
            if a[i] == "1":
                new_a= new_a + int(pow(2, i))
            i += 1
        while j < len(b):    # 将字符串b转换为十进制数
            if b[j] == "1":
                new_b = new_b + int(pow(2, j))
            j += 1
        result = str(bin(int(new_a) + int(new_b)))
        return result[2:]



if __name__ == '__main__':
    ss = Solution()
    a = "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
    b = "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
    print(ss.addBinary(a, b))
    new_a = int(a, 2)
    new_b = int(b, 2)
    print(new_a)
    print(new_b)
    print(new_a + new_b)
    print(str(bin(new_a + new_b)))
