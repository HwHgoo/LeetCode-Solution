package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit_1(prices []int) int {
	profit := 0
	minPrice := (1 << 31) ^ (^0)
	for _, p := range prices {
		profit = max(p-minPrice, profit)
		if p < minPrice {
			minPrice = p
		}
	}

	return profit
}

// @lc code=end
