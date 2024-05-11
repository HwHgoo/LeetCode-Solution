package leetcode

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	added_coins := 0
	covered_max := 0
	for _, coin := range coins {
		for covered_max+1 < coin {
			added_coins += 1
			covered_max += covered_max + 1
		}

		covered_max += coin
		if covered_max >= target {
			return added_coins
		}
	}

	for covered_max < target {
		covered_max += covered_max + 1
		added_coins += 1
	}

	return added_coins
}
