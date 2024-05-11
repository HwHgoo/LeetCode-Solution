package leetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	farthest, steps, end := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == end {
			steps++
			end = farthest
		}
	}

	return steps
}

// @lc code=end
