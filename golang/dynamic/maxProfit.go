/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
    如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
    注意你不能在买入股票前卖出股票。
 */

package main

func maxProfit(prices []int)int{
	length := len(prices)
	if length == 0{
		return 0
	}
	// 开一个大小为length的数组
	dp := make([]int, length)
	// 定义一个在遍历prices数组时当前遍历过最小的元素, 初始化prices的第一个元素
	minPrice := prices[0]
	// 从prices的第2个元素开始，分别求每天前最佳卖股票时机，计算顺序为从左到右
	for i := 1; i < length; i ++{
		sell := prices[i] - minPrice
		notSell := dp[i - 1]
		if sell > notSell{
			dp[i] = sell
		}else{
			dp[i] = notSell
		}
		if prices[i] < minPrice{
			minPrice = prices[i]
		}
	}
	return dp[length - 1]
}
