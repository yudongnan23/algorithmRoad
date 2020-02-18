"""

    给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
    如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
    注意你不能在买入股票前卖出股票。

"""


class Solution:
    def max_profit(self, prices: list) -> int:
        """

        :param prices:
        :return:
        """
        if len(prices) == 0:
            return 0
        length = len(prices)
        # 开一个大小为length的数组，且初始化都为0
        dp = [0] * length
        # 定义当前股票最小价格，初始化为股票的第0天的价格
        min_price = prices[0]
        # 从股票数组的第二个元素开始，分别求每天前的最佳卖股票时机，计算顺序为从左到右
        for i in range(1, length):
            sell = prices[i] - min_price
            not_sell = dp[i - 1]
            dp[i] = max(sell, not_sell)
            if prices[i] < min_price:
                min_price = prices[i]
        return dp[length - 1]


if __name__ == '__main__':
    price = [7, 1, 5, 3, 6, 4]
    solution = Solution()
    print(solution.max_profit(price))