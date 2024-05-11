package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func solve_wordbreakII(s string, ind int, dict []string, res *[]string, parts []string) {
	if ind > len(s) {
		return
	}

	if ind == len(s) {
		line := strings.Join(parts, " ")
		*res = append(*res, line)
		return
	}

	for _, word := range dict {
		if ind+len(word) <= len(s) && s[ind:len(word)+ind] == word {
			solve_wordbreakII(s, ind+len(word), dict, res, append(parts, word))
		}
	}
}

func wordBreak(s string, wordDict []string) []string {
	res := make([]string, 0)
	solve_wordbreakII(s, 0, wordDict, &res, nil)
	return res
}

// @lc code=end
