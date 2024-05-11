package leetcode

import "testing"

func TestHeapInsert(t *testing.T) {
	heap := NewMaxHeap([]int{1, 5, 2, 3, 6, 4, 9, 10, 8, 7})
	heap.Insert(11)
}
