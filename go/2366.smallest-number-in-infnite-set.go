package leetcode

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type SmallestInfiniteSet struct {
	dispersed  *treeset.Set
	consective int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		dispersed:  treeset.NewWithIntComparator(),
		consective: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if !this.dispersed.Empty() {
		it := this.dispersed.Iterator()
		val := it.Value().(int)
		this.dispersed.Remove(val)
		return val
	}

	defer func() {
		this.consective++
	}()

	return this.consective
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.consective {
		return
	}

	if num == this.consective-1 {
		this.consective -= 1
		return
	}

	this.dispersed.Add(num)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
