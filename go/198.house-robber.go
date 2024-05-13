package leetcode

// f(x): the max amount we can rob in houses[0, x]
// f(x) = max(f(x-1), f(x-2)+amount[x])
func rob_198(nums []int) int {
	// |-- p1: max rob in [0, cur-2] --|-- p2: max rob in [0, cur-1] --|-- cur --|
	p1, p2 := 0, 0
	for _, n := range nums {
		p1, p2 = p2, max(p1+n, p2)
	}

	return p2
}
