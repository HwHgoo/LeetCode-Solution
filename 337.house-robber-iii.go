package leetcode

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		lrob, lnrob := dfs(root.Left)
		rrob, rnrob := dfs(root.Right)
		// rob root
		rob_cur := root.Val + lnrob + rnrob
		// rob root.Left, rob root.Right, rob root.Left and Right
		nrob_cur := max(lrob, lnrob) + max(rrob, rnrob)
		return rob_cur, nrob_cur
	}

	return max(dfs(root))
}

// @lc code=end
