package leetcode

func RobII(nums []int) int {
	return rob_II(nums)
}

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve_robII(nums []int) int {
	lhs := nums[0]
	rhs := max(nums[0], nums[1])
	var tmp int
	for i := 2; i < len(nums); i++ {
		tmp = max(lhs+nums[i], rhs)
		lhs = rhs
		rhs = tmp
	}
	return rhs
}

func rob_II(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(solve_robII(nums[0:len(nums)-1]), solve_robII(nums[1:]))
}

// @lc code=end
