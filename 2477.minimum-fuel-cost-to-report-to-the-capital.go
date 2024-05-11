package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=2477 lang=golang
 *
 * [2477] Minimum Fuel Cost to Report to the Capital
 */

// @lc code=start
func minimumFuelCost(roads [][]int, seats int) int64 {
	tree := make(map[int][]int)
	for _, r := range roads {
		tree[r[0]] = append(tree[r[0]], r[1])
		tree[r[1]] = append(tree[r[1]], r[0])
	}

	var dfs func(int, int) int

	var res int64
	dfs = func(root, parent int) int {
		cur := 0
		for _, child := range tree[root] {
			if child == parent {
				continue
			}
			count := dfs(child, root)
			res += int64(math.Ceil(float64(count) / float64(seats)))
			cur += count
		}

		return cur + 1
	}

	dfs(0, -1)

	return res
}

// @lc code=end
