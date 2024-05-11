package leetcode

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	tank := 0
	for i := 0; i < len(gas); {
		wrapped := false
		k := i
		for {
			if wrapped && k == i {
				return i
			}

			if tank+gas[k] < cost[k] {
				tank = 0
				break
			} else {
				tank = tank + gas[k] - cost[k]
			}
			k++
			if k == len(gas) {
				k = 0
				wrapped = true
			}
		}
		i = k + 1
	}

	return -1
}

// @lc code=end
