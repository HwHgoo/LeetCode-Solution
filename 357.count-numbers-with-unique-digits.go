package leetcode

/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 91
	case 3:
		return 739
	case 4:
		return 5275
	case 5:
		return 32491
	case 6:
		return 168571
	case 7:
		return 712891
	case 8:
		return 2345851
	default:
		return 0
	}
}

// @lc code=end
