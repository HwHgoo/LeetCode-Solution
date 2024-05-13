package leetcode

func generate_118(numRows int) [][]int {
	ans := make([][]int, 0, numRows)
	first := []int{1}
	ans = append(ans, first)
	prev := first

	for i := 2; i <= numRows; i++ {
		cur := make([]int, len(prev)+1)
		cur[0], cur[len(prev)] = 1, 1
		for i := 1; i < len(prev); i++ {
			cur[i] = prev[i] + prev[i-1]
		}
		ans = append(ans, cur)
		prev = cur
	}

	return ans
}
