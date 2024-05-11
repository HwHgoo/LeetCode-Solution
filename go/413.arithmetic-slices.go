package leetcode

func NumberOfArithmeticSlices(nums []int) int {
	return numberOfArithmeticSlices(nums)
}

/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func my_solve_413(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	dp[1] = 2
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == diff {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 2
			diff = nums[i] - nums[i-1]
		}
	}
	count := 0
	for k := len(nums) - 1; k >= 0; {
		n := dp[k]
		if k+1 <= len(nums)-1 && dp[k+1] != 2 {
			n = dp[k+1]
		}

		if n >= 3 {
			count += (n - 2) * (n - 1) / 2
		}

		k -= n
	}
	return count

}

func easy_solve_413(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	res := 0
	diff, count := nums[1]-nums[0], 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == diff {
			count++
		} else {
			diff, count = nums[i]-nums[i-1], 0
		}
		res += count
	}

	return res
}

func numberOfArithmeticSlices(nums []int) int {
	return easy_solve_413(nums)
}

// @lc code=end
