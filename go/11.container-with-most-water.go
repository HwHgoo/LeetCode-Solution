package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

func solve_11_1(height []int) int {
	max := 0
	for l := 0; l < len(height)-1; l++ {
		for r := l + 1; r < len(height); r++ {
			if vol := min(height[l], height[r]) * (r - l); vol > max {
				max = vol
			}
		}
	}

	return max
}

func solve_11_2(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		if vol := (r - l) * min(height[l], height[r]); vol > max {
			max = vol
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

// @lc code=start
func maxArea(height []int) int {
	return solve_11_2(height)
}

// @lc code=end
