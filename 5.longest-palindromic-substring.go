package leetcode

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

func solve_5_dfs(s string) string {
	if len(s) <= 1 {
		return s
	}

	rp := len(s) - 1
	lp := 0
	count := 0
	for lp <= rp {
		if lp == rp {
			count++
			break
		} else if s[lp] == s[rp] {
			lp++
			rp--
			count += 2
		} else {
			lp = 0
			rp--
			count = 0
		}
	}
	if sub := solve_5_dfs(s[1:]); len(sub) > count {
		return sub
	}
	return s[:count]
}

// @lc code=start
func solve_5_dp(s string) string {
	pre := make([]int, len(s)+1)
	cur := make([]int, len(s)+1)

	res := s[len(s)-1:]
	pre[len(s)] = 1
	for i := len(s) - 1; i > 0; i-- {
		cur[i] = 1
		for k := i + 1; k <= len(s); k++ {
			if s[i-1] == s[k-1] && pre[k-1] == k-i-1 {
				cur[k] = 2 + k - i - 1
				if cur[k] > len(res) {
					res = s[i-1 : i-1+cur[k]]
				}
			} else {
				cur[k] = 1
			}
		}
		tmp := pre
		pre = cur
		cur = tmp
	}
	return res
}

func longestPalindrome(s string) string {
	return solve_5_dp(s)
}

// @lc code=end
