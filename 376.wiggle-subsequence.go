package leetcode

func WiggleMaxLength(nums []int) int {
	return wiggleMaxLength(nums)
}

/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	res := 1
	dire := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] > nums[i-1] {
			if dire == 1 {
				continue
			}

			dire = 1
			res += 1
		} else if nums[i] < nums[i-1] {
			if dire == 0 {
				continue
			}
			dire = 0
			res += 1
		}
	}

	return res
}

// @lc code=end
