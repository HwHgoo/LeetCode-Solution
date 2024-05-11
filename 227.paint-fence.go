package leetcode

func numWays(n int, k int) int {
	var solve func(n, k int, dup bool) int
	solve = func(n, k int, dup bool) int {
		if n == 1 {
			if dup {
				return k - 1
			} else {
				return k
			}
		}

		count := (k - 1) * solve(n-1, k, false)
		if !dup {
			count += solve(n-1, k, true)
		}

		return count
	}

	return k * solve(n-1, k, false)
}
