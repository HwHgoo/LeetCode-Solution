package leetcode

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}

	reminder := n % 3
	n -= reminder
	res := 1
	for n > 1 {
		n -= 3
		res *= 3
	}
	if reminder == 0 {
		return res
	} else if reminder == 1 {
		res = (res / 3) * 4
	} else {
		res *= reminder
	}

	return res
}

// @lc code=end
