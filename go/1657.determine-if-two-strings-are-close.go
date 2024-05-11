package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	if word1 == word2 {
		return true
	}

	count1, count2 := make([]int, 26), make([]int, 26)
	for i := range word1 {
		count1[word1[i]-'a'] += 1
		count2[word2[i]-'a'] += 1
	}

	for i := range count1 {
		if (count1[i] == 0 && count2[i] != 0) ||
			(count1[i] != 0 && count2[i] == 0) {
			return false
		}
	}

	sort.Ints(count1)
	sort.Ints(count2)
	for i := range count1 {
		if count1[i] != count2[i] {
			return false
		}
	}

	return true
}

// @lc code=end
