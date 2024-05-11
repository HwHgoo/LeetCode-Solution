package leetcode

func MaxProfix_309(prices []int) int {
	return maxProfit_with_cool_down(prices)
}

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start

func max3(x, y, z int) int {
	if x >= y && x >= z {
		return x
	}

	if y >= x && y >= z {
		return y
	}

	return z
}

const (
	CoolDown = 0
	Buy      = 1
	Sell     = 2
)

func solve309_backtrace(idx int, hold int, last_op int, prices []int) int {
	var skip, buy, sell int
	if idx >= len(prices) {
		if last_op == Buy {
			return hold
		}
		return 0
	}

	skip = solve309_backtrace(idx+1, hold, CoolDown, prices)
	if last_op != 2 && hold == -1 {
		buy = solve309_backtrace(idx+1, prices[idx], Buy, prices) - prices[idx]
	} else if hold >= 0 && prices[idx] >= hold {
		sell = prices[idx] + solve309_backtrace(idx+1, -1, Sell, prices)
	}

	return max3(skip, buy, sell)
}

func solve309_dp(dp [][]int, prices []int) int {
	for i := 1; i <= len(prices); i++ {
		dp[i][CoolDown] = max(dp[i-1][CoolDown], dp[i-1][Sell])
		dp[i][Sell] = max(dp[i-1][Buy]+prices[i-1], dp[i-1][Sell])
		dp[i][Buy] = max(dp[i-1][CoolDown]-prices[i-1], dp[i-1][Buy])
	}
	last := len(prices)
	return max3(dp[last][CoolDown], dp[last][Sell], dp[last][Buy])
}

func maxProfit_with_cool_down(prices []int) int {
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][Buy] = -prices[0]
	return solve309_dp(dp, prices)
}

// @lc code=end
