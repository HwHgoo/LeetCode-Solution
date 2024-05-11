package leetcode

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=996 lang=golang
 *
 * [996] Number of Squareful Arrays
 */

// @lc code=start
func numSquarefulPerms(nums []int) int {
	sort.Ints(nums)
	graph := make(map[int][]int)
	traveled := make(map[int]bool)
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if squareful(nums[i] + nums[j]) {
				graph[i] = append(graph[i], j)
			}
		}
	}
	res := 0
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		traveled[i] = true
		res += solve_996(nums, graph, i, traveled)
		delete(traveled, i)
	}

	return res
}

func solve_996(nums []int, graph map[int][]int, startPoint int, traveled map[int]bool) int {
	if len(traveled) == len(nums) {
		return 1
	}

	count := 0
	last := -1
	nexts := graph[startPoint]
	for i, next := range nexts {
		if traveled[next] || (i > 0 && nums[next] == last) {
			continue
		}
		traveled[next] = true
		last = nums[next]
		count += solve_996(nums, graph, next, traveled)
		delete(traveled, next)
	}
	return count
}

func squareful(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

// @lc code=end
