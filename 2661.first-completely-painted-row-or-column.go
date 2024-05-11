package leetcode

/*
 * @lc app=leetcode.cn id=2661 lang=golang
 *
 * [2661] First Completely Painted Row or Column
 */

// @lc code=start
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows := make([]int, len(mat))
	columns := make([]int, len(mat[0]))
	matmap := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matmap[mat[i][j]] = n*i + j
		}
	}

	for i, num := range arr {
		x := matmap[num] % n
		y := (matmap[num] - x) / n
		rows[y]++
		columns[x]++
		if rows[y] == n || columns[x] == m {
			return i
		}
	}

	return -1
}

// @lc code=end
