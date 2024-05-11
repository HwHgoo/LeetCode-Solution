package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	farthest := nums[0]
	start := 0
	for farthest < len(nums) && start <= farthest {
		for start <= farthest {
			dst := nums[start] + start
			start++
			if dst > farthest {
				farthest = dst
				break
			}
		}
	}

	return farthest >= len(nums)-1
}

// @lc code=end
