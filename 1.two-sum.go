package leetcode

import "sort"

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func solve_1_1(nums []int, target int) []int {
	ncp := make([]int, len(nums))
	copy(ncp, nums)
	sort.Ints(ncp)
	parts := []int{0, 0}
	for i := range ncp {
		if idx := binary_search(ncp, target-ncp[i], i+1, len(ncp)-1); idx != -1 {
			parts[0] = ncp[i]
			parts[1] = ncp[idx]
			break
		}
	}

	res := make([]int, 0, 2)
	for i := range nums {
		if nums[i] == parts[0] || nums[i] == parts[1] {
			res = append(res, i)
			if nums[i] == parts[0] {
				parts[0] = parts[1]
			} else {
				parts[1] = parts[0]
			}
		}
	}

	return res

}

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func solve_1_2(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := range nums {
		if j, exists := cache[target-nums[i]]; exists {
			return []int{i, j}
		}
		cache[nums[i]] = i
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	return solve_1_2(nums, target)
}

// @lc code=end
