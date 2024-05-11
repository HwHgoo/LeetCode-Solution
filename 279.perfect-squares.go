package leetcode

import "math"

func PerfectSquare(n int) int {
	return numSquares(n)
}

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start

func math_solve_279(n int) int {
	if is_square(n) {
		return 1
	}

	tmp := n
	for tmp%4 == 0 {
		tmp /= 4
	}

	if (tmp-7)%8 == 0 {
		return 4
	}

	for i := 1; i <= n; i++ {
		isquare := i * i
		if isquare > n {
			break
		}
		if is_square(n - isquare) {
			return 2
		}
	}
	return 3
}

func is_square(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func dp_solve_279(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	last_perfect := 1
	next_root := 2
	next_square := next_root * next_root
	for i := 2; i <= n; i++ {
		if i == next_square {
			dp[i] = 1
			next_root += 1
			next_square = next_root * next_root
			last_perfect = i
		} else if last_perfect == i-1 {
			dp[i] = 2
		} else {
			cur_root := next_root - 1
			cur_square := cur_root * cur_root
			dp[i] = 1 + dp[i-cur_square]
			for k := next_root - 2; k > 1; k-- {
				dp[i] = min(dp[i], 1+dp[i-k*k])
			}
		}
	}

	return dp[n]

}

func numSquares(n int) int {
	return math_solve_279(n)
}

// @lc code=end
