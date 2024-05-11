package leetcode

func IsUgly(n int) bool {
	return isUgly(n)
}

/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			break
		}
	}

	return n == 1
}

// @lc code=end
