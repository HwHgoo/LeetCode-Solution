package leetcode

/*
 * @lc app=leetcode.cn id=2772 lang=golang
 *
 * [2772] Apply Operations to Make All Array Elements Equal to Zero
 */

// @lc code=start
func checkArray(nums []int, k int) bool {
	cur := 0
	for i := range nums {
		if cur > nums[i] {
			return false
		}
		nums[i] -= cur
		cur += nums[i]
		if i >= k-1 {
			cur -= nums[i-k+1]
		}
	}

	return cur == 0
}

// @lc code=end

func checkArray_brute(nums []int, k int) bool {
	prefix := 0
	incre := func(x []int) bool {
		for i := 1; i < len(x); i++ {
			if x[i] < x[i-1] {
				return false
			}
		}

		return true
	}
	for {
		for prefix < len(nums) && nums[prefix] == 0 {
			prefix++
		}

		if prefix >= len(nums) {
			return true
		}

		if len(nums)-prefix < k {
			return false
		}

		if !incre(nums[prefix : prefix+k]) {
			return false
		}

		p := nums[prefix]
		for i := prefix; i < prefix+k; i++ {
			nums[i] -= p
		}
	}
}
