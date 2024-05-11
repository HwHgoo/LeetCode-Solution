package leetcode

func NthUglyNumber(n int) int {
	return nthUglyNumber(n)
}

/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	weight2 := 0
	weight3 := 0
	weight5 := 0
	var r2, r3, r5 int
	for i := 1; i < n; i++ {
		r2, r3, r5 = dp[weight2]*2, dp[weight3]*3, dp[weight5]*5
		dp[i] = min(min(r2, r3), r5)
		if dp[i] == r2 {
			weight2++
		}
		if dp[i] == r3 {
			weight3++
		}
		if dp[i] == r5 {
			weight5++
		}
	}

	return dp[n-1]
}

// @lc code=end
