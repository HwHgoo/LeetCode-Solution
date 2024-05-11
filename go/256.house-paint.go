package leetcode

func minCost(costs [][]int) int {
	pre := make([]int, 3)
	cur := make([]int, 3)
	copy(pre, costs[0])

	othertwo := func(i int) (int, int) {
		switch i {
		case 0:
			return 1, 2
		case 1:
			return 0, 2
		default:
			return 0, 1
		}
	}

	for i := 1; i < len(costs); i++ {
		for k := range cur {
			l, r := othertwo(k)
			cur[k] = costs[i][k] + min(pre[l], pre[r])
		}
		tmp := cur
		cur = pre
		pre = tmp
	}

	return min3(pre[0], pre[1], pre[2])
}
