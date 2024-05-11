package leetcode

import (
	"log"
	"strings"
)

func Solve93(s string) {
	res := restoreIpAddresses(s)
	log.Println(res)
}

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start

func easy_atoi(s string) int {
	n := 0
	for i := range s {
		n = (n * 10) + int(s[i]-'0')
	}
	return n
}

func solve_restoreip(s string, res *[]string, parts []string) {
	// fmt.Printf("solving: %s\n", s)
	if len(parts) == 4 && len(s) == 0 {
		ip := strings.Join(parts, ".")
		*res = append(*res, ip)
		return
	}

	if len(parts) == 4 || len(s) >= (4-len(parts))*3 {
		return
	}

	for i := range s {
		if i >= 3 {
			break
		}

		if i == 0 && s[i] == '0' {
			solve_restoreip(s[1:], res, append(parts, "0"))
			break
		}

		if easy_atoi(s[:i+1]) <= 255 {
			solve_restoreip(s[i+1:], res, append(parts, s[:i+1]))
		}
	}
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) <= 3 {
		return res
	}
	solve_restoreip(s, &res, nil)
	return res
}

// @lc code=end
