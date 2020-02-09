"""

    给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
    有效字符串需满足：

        左括号必须用相同类型的右括号闭合。
        左括号必须以正确的顺序闭合。
        注意空字符串可被认为是有效字符串。

"""


class Solution:
    def isValid(self, s: str) -> bool:
        """ judge the brackets of the s is valid or not

        :param s:
        :return:
        """
        is_valid = True

        open_brackets = ["(", "{", "["]
        close_brackets = [")", "}", "]"]
        stack = []
        for i in s:
            if i in open_brackets:
                stack.append(i)
                continue
            if i in close_brackets: 
                index = close_brackets.index(i)
                open_bracket = open_brackets[index]

                if not stack or stack[-1] != open_bracket:
                    is_valid = False
                    break

                stack.pop()

        if len(stack) != 0:
            is_valid = False

        return is_valid


if __name__ == '__main__':
    string = "[({}]"
    solution = Solution()
    print(solution.isValid(string))