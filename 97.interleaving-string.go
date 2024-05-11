package leetcode

/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := make([][]byte, len(s1)+1)
	for i := range dp {
		dp[i] = make([]byte, len(s2)+1)
	}

	dp[0][0] = 1
	for i := 0; i < len(s1) && s1[i] == s3[i]; i++ {
		dp[i+1][0] = 1
	}

	for i := 0; i < len(s2) && s2[i] == s3[i]; i++ {
		dp[0][i+1] = 1
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] |= dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] |= dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)] == 1
}

// @lc code=end
