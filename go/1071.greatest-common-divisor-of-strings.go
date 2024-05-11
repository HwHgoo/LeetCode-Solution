package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start

func dividedby(s string, t string) bool {
	if len(s)%len(t) != 0 {
		return false
	}

	disc := len(s) / len(t)
	return strings.Repeat(t, disc) == s
}

func gcdOfStrings(str1 string, str2 string) string {
	res := ""
	for i := min(len(str1), len(str2)); i > 0; i-- {
		t := str1[:i]
		if dividedby(str1, t) && dividedby(str2, t) {
			res = t
		}
	}

	return res
}

// @lc code=end
