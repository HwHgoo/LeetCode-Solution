package leetcode

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	// caculate left product of nums[i]
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// caculate right product of nums[i]
	right_prod := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right_prod
		right_prod *= nums[i]
	}

	return res
}

// @lc code=end
