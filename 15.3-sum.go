package leetcode

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	pc, zc, nc := 0, 0, 0
	for _, n := range nums {
		if n > 0 {
			pc++
		} else if n == 0 {
			zc++
		} else {
			nc++
		}
	}

	if zc >= 3 {
		res = append(res, []int{0, 0, 0})
	}

	nptr_max := nc - 1
	pptr_min := nc + zc

	// {0, x, y} x < 0 && y > 0
	if zc > 0 && pc > 0 && nc > 0 {
		nptr := 0
		pptr := len(nums) - 1
		for nptr < pptr && nptr <= nptr_max && pptr >= pptr_min {
			if nums[nptr]+nums[pptr] < 0 {
				nptr++
			} else if nums[nptr]+nums[pptr] > 0 {
				pptr--
			} else if pptr > 0 && nums[pptr-1] == nums[pptr] {
				pptr--
			} else if nptr < len(nums)-1 && nums[nptr+1] == nums[nptr] {
				nptr++
			} else {
				res = append(res, []int{0, nums[nptr], nums[pptr]})
				nptr++
				pptr--
			}
		}
	}

	// {x, y, z} x < 0 && y < 0 && z >0
	if pc > 0 && nc > 1 {
		for i := 0; i < nptr_max; i++ {
			if nums[i]+nums[len(nums)-1] <= 0 {
			} else if i > 0 && nums[i] == nums[i-1] {
			} else {
				nptr := i + 1
				pptr := len(nums) - 1
				target := -nums[i]
				for nptr < pptr && nptr <= nptr_max && pptr >= pptr_min {
					if nums[nptr]+nums[pptr] < target {
						nptr++
					} else if nums[nptr]+nums[pptr] > target {
						pptr--
					} else if pptr >= pptr_min && nums[pptr-1] == nums[pptr] {
						pptr--
					} else if nptr <= nptr_max && nums[nptr+1] == nums[nptr] {
						nptr++
					} else {
						res = append(res, []int{nums[i], nums[nptr], nums[pptr]})
						nptr++
						pptr--
					}
				}
			}
		}
	}

	// {x, y, z} x < 0 && y > 0 && z > 0
	if pc > 1 && nc > 0 {
		for i := len(nums) - 1; i > pptr_min; i-- {
			if nums[0]+nums[i] >= 0 {
			} else if i < len(nums)-1 && nums[i] == nums[i+1] {
			} else {
				nptr := 0
				pptr := i - 1
				target := -nums[i]
				for nptr < pptr && nptr <= nptr_max && pptr >= pptr_min {
					if nums[nptr]+nums[pptr] < target {
						nptr++
					} else if nums[nptr]+nums[pptr] > target {
						pptr--
					} else if pptr >= pptr_min && nums[pptr-1] == nums[pptr] {
						pptr--
					} else if nptr <= nptr_max && nums[nptr+1] == nums[nptr] {
						nptr++
					} else {
						res = append(res, []int{nums[i], nums[nptr], nums[pptr]})
						nptr++
						pptr--
					}
				}
			}
		}
	}

	return res

}

// @lc code=end

func threeSum_cache(nums []int) [][]int {
	res := make([][]int, 0)
	cache := make(map[string]bool)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if cache[hash(nums[i], nums[j])] {
				continue
			}
			target := -nums[i] - nums[j]

			for k := j + 1; k < len(nums); k++ {
				if nums[k] == target {
					triplet := []int{nums[i], nums[j], nums[k]}
					res = append(res, triplet)
					h1, h2, h3 := triplet_hash(triplet)
					cache[h1] = true
					cache[h2] = true
					cache[h3] = true
					continue
				}
			}
		}
	}
	return res
}

func triplet_hash(triplet []int) (string, string, string) {
	a, b, c := triplet[0], triplet[1], triplet[2]

	return hash(a, b), hash(a, c), hash(b, c)
}

func hash(x, y int) string {
	if x > y {
		tmp := x
		x = y
		y = tmp
	}
	return fmt.Sprintf("%5d%5d", x, y)
}

func threeSum_brute(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(tmp)
					id := fmt.Sprintf("%d%d", tmp[0], tmp[1])
					if m[id] {
						continue
					}
					m[id] = true
					res = append(res, tmp)
				}
			}
		}
	}
	return res
}

//[-19903, -19865, -19857, -19745, -19475, -19426, -19349, -19282, -19116, -19111, -19051, -18931, -18908, -18636, -18563, -18484, -18250, -18215, -17915, -17869, -17537, -17516, -17474, -17410, -17351, -17228, -17174, -17077, -17037, -17037, -16983, -16917, -16907, -16805, -16758, -16682, -16500, -16374, -16364, -16308, -16252, -16147, -16144, -16085, -16012, -15895, -15844, -15840, -15737, -15715, -15698, -15678, -15442, -15298, -15236, -14959, -14724, -14660, -14485, -14395, -14345, -14278, -14267, -13909, -13581, -13496, -13418, -13267, -13261, -13164, -13103, -12946, -12869, -12840, -12786, -12659, -12651, -12509, -12406, -12367, -12341, -12311, -12259, -11730, -11499, -11293, -11141, -11139, -11014, -10989, -10959, -10922, -10806, -10630, -10551, -10462, -10323, -10288, -10278, -10109, -9985, -9862, -9827, -9581, -9256, -9077, -8998, -8977, -8838, -8716, -8548, -8509, -7870, -7771, -7736, -7677, -7314, -7305, -7245, -7237, -7219, -7066, -7004, -6699, -6674, -6647, -6647, -6600, -6496, -6341, -6328, -6146, -6137, -5856, -5853, -5787, -5595, -5586, -5412, -5372, -5220, -5171, -4988, -4970, -4933, -4895, -4794, -4781, -4779, -4641, -4621, -4589, -4392, -4239, -3902, -3680, -3541, -3464, -3455, -3454, -3255, -3196, -2986, -2940, -2916, -2898, -2625, -2406, -2402, -2397, -2162, -2110, -1900, -1182, -1108, -1080, -1076, -893, -869, -754, -637, -487, -387, -363, -302, -33, 340, 1965, 2149, 2441, 4024, 5764, 5813, 5825, 5895, 6141, 6335, 6476, 7082, 7186, 7209, 7558, 8712, 8883, 8967, 9328, 9371, 9540, 9664, 9914, 10517, 10693, 10779, 11063, 11129, 11239, 11683, 13330, 14573, 14685, 14783, 15157, 15758, 16172, 16245, 17215, 18575, 18807, 18882, 19191, 19322, 19908, 20106, 20456, 20572, 21249, 21725, 21796, 21916, 22589, 23257, 23842, 24736, 25187, 25195, 25413, 25667, 25934, 26201, 26526, 26824, 27367, 27629, 27761, 28396, 28572, 28691, 29311, 31893, 32165, 32239, 32353, 34790, 35153, 36010, 36203, 36757, 37164, 37261, 38101, 38219, 39028, 39226, 39326, 39875, 40075, 40199, 40761, 40841, 41498, 41727, 41765, 42036, 42640, 42776, 43183, 43347, 43914, 44202, 44244, 44593, 45147, 45409, 45573, 45706, 45993, 46748, 46754, 48281, 48293, 49551, 50627, 50629, 51071, 51126, 51474, 53356, 54106, 54286, 54627, 55100, 56469, 57201, 57607, 58325, 58727, 59087, 59670, 60803, 60870, 61188, 61873, 62860, 63060, 64009, 65391, 65609, 65940, 66426, 66860, 66874, 66961, 67403, 67797, 68009, 68203, 69639, 69647, 69766, 69990, 70084, 70642, 70902, 71493, 71542, 71687, 71753, 71995, 72248, 72340, 73606, 74193, 74336, 74475, 74769, 74993, 76144, 76178, 76247, 76445, 76857, 77121, 78200, 78203, 78346, 78529, 78644, 79113, 80250, 80674, 80885, 80989, 81994, 82184, 82792, 82814, 83567, 83706, 83892, 85292, 85304, 85568, 86579, 86805, 86851, 88353, 88398, 88651, 88914, 88976, 89127, 89469, 90023, 90445, 91401, 91988, 92543, 92758, 93791, 94604, 95825, 95980, 96301, 97971, 98205, 98701, 99083, 99144, 99341, 99346, 99368]
