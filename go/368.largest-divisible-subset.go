package leetcode

import (
	"sort"
)

func LargestDivisibleSubset(nums []int) []int {
	return largestDivisibleSubset(nums)
}

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start
// Time: O(n^2)
// Space: O(n^2)
func solve_368_n2(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	sort.Ints(nums)

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = []int{nums[i]}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j])+1 > len(dp[i]) {
				dp[i] = append([]int{nums[i]}, dp[j]...)
			}
		}
	}

	maxidx := 0
	for i := range dp {
		if len(dp[i]) > len(dp[maxidx]) {
			maxidx = i
		}
	}

	return dp[maxidx]

}

// space O(n)
// time O(n^2)
func solve_368_n(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	max_val := 0
	max_idx := 0
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > dp[max_idx] {
					max_val = nums[i]
					max_idx = i
				}
			}
		}
	}

	max_size := dp[max_idx]
	res := make([]int, 0, max_size)
	for i := max_idx; i >= 0 && max_size > 0; i-- {
		if dp[i] == max_size && max_val%nums[i] == 0 {
			res = append(res, nums[i])
			max_val = nums[i]
			max_size--
		}
	}

	return res
}

func largestDivisibleSubset(nums []int) []int {
	return solve_368_n(nums)
}

// @lc code=end
