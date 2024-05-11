package leetcode

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0, factorial(len(nums)))
	backtrace(nums, &res, make([]int, 0, len(nums)))
	return res
}

func backtrace(nums []int, res *[][]int, cur []int) {
	if len(nums) == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := range nums {
		curlen := len(cur)
		cur = append(cur, nums[i])
		left := make([]int, len(nums[:i])+len(nums[i+1:]))
		copy(left, nums[:i])
		copy(left[i:], nums[i+1:])
		backtrace(left, res, cur)
		cur = cur[:curlen]
	}
}

func factorial(n int) int {
	res := 1
	for n > 1 {
		res *= n
		n--
	}
	return res
}

// @lc code=end
