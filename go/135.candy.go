package leetcode

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	left, right := make([]int, len(ratings)), make([]int, len(ratings))
	for i := range ratings {
		left[i], right[i] = 1, 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	sum := 0
	for i := range ratings {
		sum += max(left[i], right[i])
	}

	return sum
}

// @lc code=end
