package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] Valid Number
 */

// @lc code=start
func isNumber(s string) bool {
	const (
		valid = iota
		invalid
		empty
	)
	s = strings.ToLower(s)
	integer := func(ns string) int {
		if len(ns) == 0 {
			return empty
		}

		nstart := 0
		for ns[nstart] == '+' || ns[nstart] == '-' {
			nstart++
		}
		if nstart >= len(ns) {
			return empty
		}

		for i := nstart; i < len(ns); i++ {
			if ns[i] < '0' || ns[i] > '9' {
				return invalid
			}
		}

		return valid
	}

	valid_decimal := func(ds string) bool {
		parts := strings.Split(ds, ".")
		if len(parts) > 2 {
			return false
		}
		if len(parts) == 1 {
			return integer(ds) == valid
		}
		front, back := integer(parts[0]), integer(parts[1])
		if front == invalid || back == invalid {
			return false
		}

		return front != empty || back != empty
	}

	parts := strings.Split(s, "e")
	if len(parts) > 2 {
		return false
	}

	if len(parts) == 1 {
		return valid_decimal(s)
	}

	return valid_decimal(parts[0]) && integer(parts[1]) == valid
}

// @lc code=end
