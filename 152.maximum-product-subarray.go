package leetcode

func MaxProduct(nums []int) int {
	return maxProduct(nums)
}

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start

func min3(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}

	if y <= x && y <= z {
		return y
	}
	return z
}

func maxProduct(nums []int) int {
	maxdp := nums[0]
	mindp := nums[0]
	res := nums[0]

	var tmp int

	for i := 1; i < len(nums); i++ {
		// maxdp
		tmp = max3(maxdp*nums[i], mindp*nums[i], nums[i])
		mindp = min3(maxdp*nums[i], mindp*nums[i], nums[i])
		maxdp = tmp
		if maxdp > res {
			res = maxdp
		}
	}

	return res
}

// @lc code=end
