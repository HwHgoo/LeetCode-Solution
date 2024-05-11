package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge_interval(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])
	cur := 0
	for _, interval := range intervals[1:] {
		if interval[0] <= merged[cur][1] {
			merged[cur][1] = max(interval[1], merged[cur][1])
		} else {
			merged = append(merged, interval)
			cur++
		}
	}

	return merged
}

// @lc code=end

func solve56_with_bitmap(intervals [][]int) [][]int {
	max := 0
	bm := NewBitmap(20001)
	for _, interval := range intervals {
		start, end := interval[0]*2, interval[1]*2
		if end > max {
			max = end
		}
		for i := start; i <= end; i++ {
			bm.Set(i)
		}
	}

	res := make([][]int, 0)
	iter := bm.Iterator()
	for iter.cur <= max {
		start := iter.NextSet() / 2
		end := (iter.NextUnset() - 1) / 2
		res = append(res, []int{start, end})
	}
	return res

}
