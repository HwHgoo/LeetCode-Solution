package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	digit_count := func(num int) int {
		if num == 0 {
			return 1
		}
		count := 0
		for num != 0 {
			count++
			num /= 10
		}
		return count
	}
	// avoid use 64-bits (u)int
	easy_itoa := func(num int) string {
		if num == 0 {
			return "0"
		}
		sign := 1
		digits := digit_count(num) + 1
		buf := make([]byte, digits)
		if num < 0 {
			buf[0] = byte('-')
			sign = -1
		} else {
			buf[0] = byte('+')
		}

		i := digits - 1
		for num != 0 {
			buf[i] = byte((num%10)*sign) + 48
			i--
			num /= 10
		}
		return string(buf)
	}
	easy_atoi := func(str string) int {
		num, base := 0, 1
		if str[0] == '-' {
			base = -1
		}
		str = str[1:]
		for i := len(str) - 1; i >= 0; i-- {
			num += int((str[i] - 48)) * base
			base *= 10
		}
		return num
	}

	reverse_numstr := func(str string) string {
		builder := strings.Builder{}
		builder.Grow(len(str))
		builder.WriteByte(str[0])
		for i := len(str) - 1; i >= 1; i-- {
			builder.WriteByte(str[i])
		}
		return builder.String()
	}

	const (
		INT_MIN = -1 << 31
		INT_MAX = (1 << 31) - 1
	)
	MIN_STR, MAX_STR := easy_itoa(INT_MIN), easy_itoa(INT_MAX)

	xstr := easy_itoa(x)
	reversed := reverse_numstr(xstr)
	if len(reversed) == len(MIN_STR) && (reversed > MAX_STR || reversed < MIN_STR) {
		return 0
	}

	return easy_atoi(reversed)
}

// @lc code=end
