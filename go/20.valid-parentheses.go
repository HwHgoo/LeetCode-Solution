package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	top := 1
	stack := make([]rune, len(s)+1)
	for _, p := range s {
		if diff := p - stack[top-1]; diff > 0 && diff <= 2 {
			top--
		} else {
			stack[top] = p
			top++
		}
	}

	return top == 1

}

// @lc code=end

/*
	if len(s)&1 != 0 {
		return false
	}

	top := 1
	stack_size := len(s)>>1 + 2
	stack := make([]rune, stack_size)
	for _, p := range s {
		if top >= stack_size {
			return false
		}

		if diff := p - stack[top-1]; diff > 0 && diff <= 2 {
			top--
		} else {
			stack[top] = p
			top++
		}
	}

	return top == 1

*/
