package leetcode

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start

func wordBreak_I(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}
	dp[0] = true
	max_word_len := 0
	for i := range wordDict {
		max_word_len = max(max_word_len, len(wordDict[i]))
	}

	// i should start from 1
	// so use i+1
	for i := range s {
		for j := i; j >= max(i+1-max_word_len, 0); j-- {
			if dp[j] && words[s[j:i+1]] {
				dp[i+1] = true
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end
