package leetcode

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	res := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			lbrac_idx := i - dp[i-1] - 1
			if lbrac_idx >= 0 && s[lbrac_idx] == '(' {
				dp[i] = dp[i-1] + 2
				if lbrac_idx-1 >= 0 {
					dp[i] += dp[lbrac_idx-1]
				}
				res = max(dp[i], res)
			}
		}
	}

	return res
}

// @lc code=end
