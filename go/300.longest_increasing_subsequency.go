package leetcode

import "sort"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	ans := 0
	for i := 1; i <= len(nums); i++ {
		dp[i] = 1
		for j := 1; j <= i-1; j++ {
			if nums[i-j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[i-j]+1)
			}
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func lengthOfLIS_b(nums []int) int {
	// ascending sequency
	// ascSeq[i] : tail element of ascending subsequency with (i+1) elements
	ascSeq := []int{}
	for _, n := range nums {
		idx := sort.SearchInts(ascSeq, n)
		if idx == len(ascSeq) {
			ascSeq = append(ascSeq, n)
		} else {
			ascSeq[idx] = n
		}
	}

	return len(ascSeq)
}
