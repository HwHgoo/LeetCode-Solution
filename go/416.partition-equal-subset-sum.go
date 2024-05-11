package leetcode

func CanPartition(nums []int) bool {
	return canPartition(nums)
}

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start

// space O(n^2)
func solv416_1(nums []int) bool {
	sum := 0
	max := 0
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if sum%2 != 0 {
		return false
	}

	half := sum / 2
	if max > half {
		return false
	}
	if max == half {
		return true
	}

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		for j := 1; j < half+1; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][half]
}

// space O(n)
func solv416_2(nums []int, half int) bool {
	var tmp []bool
	dp_pre := make([]bool, half+1)
	dp_cur := make([]bool, half+1)
	dp_pre[0], dp_pre[nums[0]], dp_cur[0] = true, true, true

	for i := 1; i < len(nums); i++ {
		for j := 1; j < half+1; j++ {
			if j >= nums[i] {
				dp_cur[j] = dp_pre[j] || dp_pre[j-nums[i]]
			} else {
				dp_cur[j] = dp_pre[j]
			}
		}

		tmp = dp_cur
		dp_cur = dp_pre
		dp_pre = tmp
	}

	return dp_pre[half]
}


// optimize space use with bitmap
func solv416_3(nums []int, half int) bool {
	pre := NewBitmap(half + 1)
	cur := NewBitmap(half + 1)
	var tmp *Bitmap
	pre.Set(0)
	pre.Set(nums[0])
	cur.Set(0)
	for i := 1; i < len(nums); i++ {
		for j := 1; j < half+1; j++ {
			if j >= nums[i] {
				if pre.Get(j) || pre.Get(j-nums[i]) {
					cur.Set(j)
				}
			} else {
				if pre.Get(j) {
					cur.Set(j)
				}
			}
		}

		tmp = cur
		cur = pre
		pre = tmp
	}

	return pre.Get(half)
}

func solv416_4(nums []int, half int) bool {
	bm := NewBitmap(half + 1)
	bm.Set(0)
	for i := range nums {
		bm.Set(nums[i])
	}
	return false
}

func canPartition(nums []int) bool {
	sum := 0
	max := 0
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if sum%2 != 0 {
		return false
	}

	half := sum / 2
	if max > half {
		return false
	}
	if max == half {
		return true
	}
	return solv416_3(nums, half)
}

// @lc code=end
