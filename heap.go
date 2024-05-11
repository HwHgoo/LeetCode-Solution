package leetcode

import "sort"

type number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~uint | ~float32 | ~float64
}

type Heap[T number] interface {
	Insert(T)
	Delete(T)

	Top() T
}

type MaxHeap[T number] struct {
	h []T // starts from index 1

	size int
}

func NewMaxHeap[T number](values []T) *MaxHeap[T] {
	// create a max heap from the values
	h := make([]T, len(values)+1)
	copy(h[1:], values)
	sort.Slice(h[1:], func(i, j int) bool { return h[i+1] > h[j+1] })
	return &MaxHeap[T]{h: h, size: len(values)}
}

func (heap *MaxHeap[T]) floatUp(i int) {
	// move the value at index i up to its correct position in the heap
	v := heap.h[i]
	p := heap.parant(i)
	pv := heap.h[p]
	for i > 1 && v > pv {
		heap.h[i], heap.h[p] = pv, v
		i = p
		v = heap.h[i]
		p = heap.parant(i)
		pv = heap.h[p]
	}
}

func (heap *MaxHeap[T]) parant(i int) int {
	// return the parent index of i
	// left child
	if i%2 == 1 {
		return i / 2
	}

	return (i - 1) / 2
}

func (heap *MaxHeap[T]) leftChild(i int) int {
	// return the left child index of i
	return 2 * i
}

func (heap *MaxHeap[T]) rightChild(i int) int {
	// return the right child index of i
	return 2*i + 1
}

func (heap *MaxHeap[T]) Insert(v T) {
	// insert a value into the heap
	if heap.size+1 == len(heap.h) {
		// heap is full, double the size
		heap.h = append(heap.h, v)
	} else {
		heap.h[heap.size+1] = v
	}
	heap.size++
	heap.floatUp(heap.size)
}

func (heap *MaxHeap[T]) Delete() {
	// delete a value from the heap
}

func (heap *MaxHeap[T]) Top() T {
	// return the top value of the heap
	if heap.size == 0 {
		return 0
	}

	return heap.h[1]
}
