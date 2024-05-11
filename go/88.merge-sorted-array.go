package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	merged := make([]int, 0, m+n)

	var i, k int
	for i, k = 0, 0; i < m && k < n; {
		if nums1[i] < nums2[k] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[k])
			k++
		}
	}
	left := append(nums1[i:m], nums2[k:]...)
	merged = append(merged, left...)
	copy(nums1, merged)
}

// @lc code=end
