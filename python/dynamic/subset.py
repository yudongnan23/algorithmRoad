"""

    给定一个数组arr，和一个目标数S，从数组随机挑出任意个数，若能使得这几个数的和等于目标数S，返回True，否则返回False

"""


from copy import deepcopy


class Solution:
    def recursive_subset(self, arr: list, S: int) -> bool:
        """ judge the problem by recursive

        :param arr:
        :param S:
        :return:
        """
        cur_index = len(arr) - 1
        result = self.recursive_judge(arr, cur_index, S)
        return result

    def recursive_judge(self, arr: list, cur_index: int, S: int) -> bool:
        print(S, cur_index)
        if S == 0 or arr[cur_index] == S:
            return True
        elif cur_index == 0:
            return arr[0] == S
        elif arr[cur_index] > S:
            return self.recursive_judge(arr, cur_index - 1, S)
        else:
            select_cur_result = self.recursive_judge(arr, cur_index - 1, S - arr[cur_index])
            not_select_result = self.recursive_judge(arr, cur_index - 1, S)
            return select_cur_result or not_select_result

    def dp_subset(self, arr: list, S: int) -> bool:
        """ judge the problem by dynamic programing

        :param arr:
        :param S:
        :return:
        """
        arr_length = len(arr)
        subset = []
        sub = [False] * (S + 1)

        for i in range(arr_length):
            cp_sub = deepcopy(sub)
            cp_sub[0] = True
            if i == 0:
                cp_sub[arr[0]] = True
            subset.append(cp_sub)

        for i in range(1, arr_length):
            for s in range(1, S + 1):
                if arr[i] > s:
                    subset[i][s] = subset[i - 1][s]
                else:
                    select_result = subset[i - 1][s - arr[i]]
                    not_select_result = subset[i - 1][s]
                    subset[i][s] = select_result or not_select_result

        return subset[arr_length - 1][S]


if __name__ == '__main__':
        arr = [3, 34, 4, 12, 5, 2]
        Sum = 9
        solution = Solution()
        print(solution.recursive_subset(arr, Sum))
        print(solution.dp_subset(arr, Sum))
