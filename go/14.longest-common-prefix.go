package leetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	prefix_len := func(prefix, str string) int {
		for i := 0; i < len(str) && i < len(prefix); i++ {
			if prefix[i] != str[i] {
				return i
			}
		}
		return max(len(prefix), len(str))
	}

	for i := 1; i < len(strs); i++ {
		if prefix == "" {
			break
		}
		rear := prefix_len(prefix, strs[i])
		prefix = prefix[:rear]
	}
	return prefix
}

// @lc code=end
