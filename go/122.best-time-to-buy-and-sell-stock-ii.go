package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

func solve122(prices []int) int {
	return maxProfit(prices)
}

// @lc code=start
func maxProfit(prices []int) int {
	buy, hold, sell := -prices[0], 0, 0
	var tmp_buy, tmp_hold, tmp_sell int
	for i := 1; i < len(prices); i++ {
		tmp_hold = max(hold, sell)
		tmp_sell = prices[i] + buy
		tmp_buy = max(-prices[i]+max(hold, sell), buy)

		buy, hold, sell = tmp_buy, tmp_hold, tmp_sell
	}

	return max(hold, sell)
}

// @lc code=end
