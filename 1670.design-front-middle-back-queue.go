package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=1670 lang=golang
 *
 * [1670] Design Front Middle Back Queue
 */

// @lc code=start
type FrontMiddleBackQueue struct {
	front  *QueueNode
	middle *QueueNode
	length int
}

type QueueNode struct {
	Prev *QueueNode
	Next *QueueNode
	Val  int
}

func NewFrontMiddleBackQueue() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	node := &QueueNode{Val: val}
	if this.front == nil {
		this.front = node
		this.front.Next = this.front
		this.front.Prev = this.front
		this.middle = this.front
	} else {
		rear := this.front.Prev
		rear.Next = node
		node.Prev = rear
		node.Next = this.front
		this.front.Prev = node
		this.front = node

	}
	this.length++
	// move mid
	if this.length%2 == 0 {
		this.middle = this.middle.Prev
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	node := &QueueNode{Val: val}
	if this.front == nil {
		this.PushFront(val)
		return
	}
	var insertn *QueueNode
	if this.length%2 == 0 {
		insertn = this.middle.Next
	} else {
		insertn = this.middle
	}

	node.Next = insertn
	node.Prev = insertn.Prev
	insertn.Prev = node
	node.Prev.Next = node

	this.middle = node
	if insertn == this.front {
		this.front = node
	}
	this.length++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	if this.front == nil {
		this.PushFront(val)
		return
	}
	node := &QueueNode{Val: val}
	node.Next = this.front
	node.Prev = this.front.Prev
	node.Prev.Next = node
	this.front.Prev = node
	this.length++

	if this.length%2 == 1 {
		this.middle = this.middle.Next
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.length == 0 {
		return -1
	}

	val := this.front.Val
	this.length--
	if this.length == 0 {
		this.front = nil
		this.middle = nil
		return val
	}

	newFront := this.front.Next
	newFront.Prev = this.front.Prev
	rear := this.front.Prev
	rear.Next = this.front.Next
	this.front = newFront
	if this.length%2 != 0 {
		this.middle = this.middle.Next
	}
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.length == 0 {
		return -1
	}
	this.length--
	val := this.middle.Val
	if this.length == 0 {
		this.front = nil
		this.middle = nil
		return val
	}

	var newmid *QueueNode
	if this.length%2 == 0 {
		newmid = this.middle.Prev
		newmid.Next = this.middle.Next
		this.middle.Next.Prev = newmid
	} else {
		newmid = this.middle.Next
		newmid.Prev = this.middle.Prev
		this.middle.Prev.Next = newmid
	}

	this.middle = newmid
	if this.length == 1 {
		this.front = this.middle
	}
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.length == 0 {
		return -1
	}
	rear := this.front.Prev
	val := rear.Val
	this.length--
	if this.length == 0 {
		this.front = nil
		this.middle = nil
		return val
	}

	this.front.Prev = rear.Prev
	rear.Prev.Next = this.front
	if this.length%2 == 0 {
		this.middle = this.middle.Prev
	}
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
// @lc code=end

func print_queue(q *FrontMiddleBackQueue) {
	if q.front == nil {
		fmt.Println("nil")
		return
	}

	res := fmt.Sprintf("%d->", q.front.Val)
	node := q.front.Next
	for node != q.front {
		res += fmt.Sprintf("%d->", node.Val)
		node = node.Next
	}
	res += "head\n"
	print(res)
}
