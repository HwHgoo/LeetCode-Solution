package leetcode

func Solve1143(t1, t2 string) int {
	return longestCommonSubsequence(t1, t2)
}

/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, 0, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp = append(dp, make([]int, len(text2)+1))
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end
