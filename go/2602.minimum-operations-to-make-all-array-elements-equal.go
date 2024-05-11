package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2602 lang=golang
 *
 * [2602] Minimum Operations to Make All Array Elements Equal
 */

// @lc code=start
func minOperations(nums []int, queries []int) []int64 {
	answers := make([]int64, len(queries))
	sort.Ints(nums)
	// find the smallest index that nums[index] >= target
	var bs func(int, int, int) int
	bs = func(target, start, end int) int {
		if target < nums[start] {
			return start
		}

		if target > nums[end] {
			return end + 1
		}

		mid := start + (end-start)/2
		nmid := nums[mid]
		if nmid == target || nmid > target && nums[mid-1] < target {
			return mid
		} else if nmid < target {
			return bs(target, mid+1, end)
		} else if nmid > target {
			return bs(target, start, mid-1)
		}
		return -1
	}

	prefix := make([]int64, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefix[i] = prefix[i-1] + int64(nums[i-1])
	}

	for i, query := range queries {
		idx := bs(query, 0, len(nums)-1)
		less_count := idx
		greater_count := len(nums) - idx
		answers[i] = int64(less_count*query) - prefix[idx] + prefix[len(nums)] - prefix[idx] - int64(greater_count*query)
	}

	return answers
}

// @lc code=end
