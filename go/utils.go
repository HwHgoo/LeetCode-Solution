package leetcode

type SignedInt interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8
}

type UnsignedInt interface {
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8
}

type Int interface {
	SignedInt | UnsignedInt
}

// ints must be sort accending
// return index if target found
// return -1 if not found
func binary_search[T Int](nums []T, target T, start, end int) int {
	if start > end {
		return -1
	}

	mid := (end-start)/2 + start
	cur := nums[mid]
	if cur == target {
		return mid
	} else if nums[mid] > target {
		return binary_search(nums, target, start, mid-1)
	} else {
		return binary_search(nums, target, start+1, end)
	}
}

// nums must not be empty
func min[T int | uint](nums ...T) T {
	r := nums[0]
	for i := range nums {
		if nums[i] < r {
			r = nums[i]
		}
	}
	return r
}
