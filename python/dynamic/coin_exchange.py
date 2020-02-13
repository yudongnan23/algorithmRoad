"""
    给定一组硬币和需要的钱的总数，组合给定的硬币得到钱的总数，并使得硬币数最少。

"""

class Solution:
    def coin_exchange(self, coins: list, money: int) -> int:
        if money < 0:
            return None
        f = [0] * ( money + 1)
        length = len(coins)
        MAX_VALUE = 9999
        for i in range(1, money + 1):
            min_num = MAX_VALUE
            for j in coins:
                if i >= j and min_num > f[i - j] + 1:
                    min_num = f[i - j] + 1
            f[i] = min_num
        print(f)
        if f[money] == MAX_VALUE:
            return -1
        return f[money]


if __name__ == '__main__':
    coin = [2, 5, 7]
    money = 27
    solution = Solution()
    print(solution.coin_exchange(coin, money))