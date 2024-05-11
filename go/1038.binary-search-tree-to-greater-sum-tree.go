package leetcode

/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] Binary Search Tree to Greater Sum Tree
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func bstToGst(root *TreeNode) *TreeNode {
	right_sum := 0
	sum_travel(root, &right_sum)
	return root
}

func sum_travel(root *TreeNode, right_sum *int) {
	if root.Right != nil {
		sum_travel(root.Right, right_sum)
	}

	root.Val += *right_sum
	*right_sum = root.Val

	if root.Left != nil {
		sum_travel(root.Left, right_sum)
	}
}

// @lc code=end
