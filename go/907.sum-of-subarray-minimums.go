package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] Sum of Subarray Minimums
 */

// @lc code=start
func sumSubarrayMins(arr []int) int {
	stack := make([]int, len(arr))
	top := -1
	var res uint64
	const M uint64 = 1e+9 + 7

	var l, r int
	for i := range arr {
		for top >= 0 && arr[i] < arr[stack[top]] {
			if top > 0 {
				l = stack[top] - stack[top-1]
			} else {
				l = stack[top] + 1
			}
			r = i - stack[top]
			res = (uint64(l*r*arr[stack[top]]) + res) % M
			top--
		}

		// push
		top++
		stack[top] = i
	}

	for top >= 0 {
		if top > 0 {
			l = stack[top] - stack[top-1]
		} else {
			l = stack[top] + 1
		}
		r = len(arr) - stack[top]
		res = (uint64(l*r*arr[stack[top]]) + res) % M
		top--
	}

	return int(res)
}

// @lc code=end

func solve_907_brute(arr []int) int {
	res := 0
	m := int(math.Pow10(9)) + 7
	for i := range arr {
		min := arr[i]
		for l := 1; l <= len(arr)-i; l++ {
			if min > arr[i+l-1] {
				min = arr[i+l-1]
			}
			res += min
			if res > m {
				res %= m
			}
		}
	}

	return res
}

func solve_907_brute_opt(arr []int) int {
	var res uint64
	mod := uint64(math.Pow10(9)) + 7
	var l, r int
	for i := range arr {
		for l = i - 1; l >= 0 && arr[i] < arr[l]; l-- {
		}
		for r = i + 1; r < len(arr) && arr[i] <= arr[r]; r++ {
		}
		res = (uint64((i-l)*(r-i))*uint64(arr[i]) + res) % mod
	}

	return int(res)

}
