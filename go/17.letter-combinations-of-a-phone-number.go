package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	digitsmap := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var prod []string
	for _, digit := range digits {
		prod = slice_multi(prod, digitsmap[rune(digit)])
	}

	return prod
}

func slice_multi(l, r []string) []string {
	res := make([]string, 0, len(l)*len(r))
	if len(l) == 0 || len(r) == 0 {
		return append(l, r...)
	}

	for _, ls := range l {
		for _, rs := range r {
			res = append(res, ls+rs)
		}
	}
	return res
}

// @lc code=end
