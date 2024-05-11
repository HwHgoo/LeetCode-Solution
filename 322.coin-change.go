package leetcode

func CoinChange322(coins []int, amount int) int {
	return coinChange(coins, amount)
}

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// we cannot use math.Inf
	const MAX uint = ^uint(0) - 1
	dp := make([]uint, amount+1)
	for i := range dp {
		dp[i] = MAX
	}

	dp[0] = 0
	for i := uint(1); i <= uint(amount); i++ {
		for _, coin := range coins {
			if uint(coin) > i {
				continue
			}
			// convert to uint to compare
			dp[i] = min(1+dp[i-uint(coin)], dp[i])
		}
	}

	if dp[amount] == MAX {
		return -1
	}

	return int(dp[amount])
}

// @lc code=end
