package leetcode

func climbStairs(n int) int {
	return dp(n)
}

func dp(n int) int {
	p1, p2 := 1, 1
	total := 1
	for i := 2; i <= n; i++ {
		total = p1 + p2
		p1, p2 = p2, total
	}

	return total
}

func climb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if k, exists := memo[n]; exists {
		return k
	}

	total := climb(n-1, memo) + climb(n-2, memo)
	memo[n] = total
	return total
}
