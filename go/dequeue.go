package leetcode

import (
	"math/rand"
	"time"
)

// double linked list
type DequeueNode[T any] struct {
	val  T
	prev *DequeueNode[T]
	next *DequeueNode[T]
}

type Dequeue[T any] struct {
	head *DequeueNode[T]
	tail *DequeueNode[T]
	size int
}

func NewDequeue[T any]() *Dequeue[T] {
	head := &DequeueNode[T]{}
	tail := &DequeueNode[T]{}
	head.next = tail
	tail.prev = head
	return &Dequeue[T]{head: head, tail: tail, size: 0}
}

func (d *Dequeue[T]) PushNodeBack(node *DequeueNode[T]) {
	d.tail.prev.next = node
	node.prev = d.tail.prev
	node.next = d.tail
	d.tail.prev = node
	d.size++
}

func (d *Dequeue[T]) RemoveNode(node *DequeueNode[T]) {
	if node == nil {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	d.size--
}

func (d *Dequeue[T]) LastNode() *DequeueNode[T] {
	if d.size == 0 {
		return nil
	}
	return d.tail.prev
}

func (d *Dequeue[T]) FirstNode() *DequeueNode[T] {
	if d.size == 0 {
		return nil
	}
	return d.head.next
}

func (d *Dequeue[T]) PushBack(val T) {
	node := &DequeueNode[T]{val: val}
	d.PushNodeBack(node)
}

func (d *Dequeue[T]) Size() int {
	return d.size
}

type aaa struct {
	q *Dequeue[struct{}]
}

// generate palindrome numbers, split with ','

func generatePalindromes(count int) []int {
	if count == 0 {
		return nil
	}
	buf := make([]int, count)
	for l, r := 0, count-1; l <= r; {
		rand.NewSource(time.Now().UnixMicro())
		n := rand.Intn(10)
		buf[l] = n
		buf[r] = n

		l++
		r--
	}

	return buf
}
