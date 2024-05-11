package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	thousands := num / 1000
	hundreds := (num - 1000*thousands) / 100
	tens := (num - 1000*thousands - 100*hundreds) / 10
	ones := num % 10
	roman := ""

	roman += strings.Repeat("M", thousands)
	if hundreds == 9 {
		roman += "CM"
	} else if hundreds == 4 {
		roman += "CD"
	} else {
		if hundreds >= 5 {
			roman += "D"
			hundreds -= 5
		}
		roman += strings.Repeat("C", hundreds)
	}

	if tens == 9 {
		roman += "XC"
	} else if tens == 4 {
		roman += "XL"
	} else {
		if tens >= 5 {
			roman += "L"
			tens -= 5
		}
		roman += strings.Repeat("X", tens)
	}

	if ones == 9 {
		roman += "IX"
	} else if ones == 4 {
		roman += "IV"
	} else {
		if ones >= 5 {
			roman += "V"
			ones -= 5
		}
		roman += strings.Repeat("I", ones)
	}

	return roman
}

// @lc code=end
