package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	next := func(cur []int, dep int) []int {
		aval := 0x1ff
		for i := 0; i < dep; i++ {
			aval &= ^(1 << cur[i])
			aval &= ^(1 << (cur[i] + dep - i))
			if diff := cur[i] - (dep - i); diff >= 0 {
				aval &= ^(1 << diff)
			}
		}

		res := make([]int, 0)
		for i := 0; i < n; i++ {
			if aval&(1<<i) != 0 {
				res = append(res, i)
			}
		}
		return res
	}

	var bt func(dep int, cur []int, res *[][]int)
	cur := make([]int, n)
	res := make([][]int, 0)
	bt = func(dep int, cur []int, res *[][]int) {
		if dep == n {
			t := make([]int, n)
			copy(t, cur)
			(*res) = append((*res), t)
			return
		}
		for _, k := range next(cur, dep) {
			cur[dep] = k
			bt(dep+1, cur, res)
		}
	}
	for i := 0; i < n; i++ {
		cur[0] = i
		bt(1, cur, &res)
	}

	result := func(res [][]int) [][]string {
		r := make([][]string, len(res))
		for i := range res {
			for _, k := range res[i] {
				r[i] = append(r[i], strings.Repeat(".", k)+"Q"+strings.Repeat(".", n-k-1))
			}
		}

		return r
	}

	return result(res)
}

// @lc code=end
