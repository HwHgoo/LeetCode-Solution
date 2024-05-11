package leetcode

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeFromSlice(nums ...int) *ListNode {
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	cur := head
	for _, num := range nums[1:] {
		node := &ListNode{
			Val:  num,
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	return head
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for node := head; node != nil; {
		length++
		node = node.Next
	}

	target_idx := length - n
	if target_idx == 0 {
		return head.Next
	}

	node := head
	for i := 0; i < target_idx-1; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next
	return head
}

// @lc code=end
