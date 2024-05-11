package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

// @lc code=end

func addTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		cur.Next = &ListNode{}
		cur = cur.Next

		cur.Val = l1.Val + l2.Val + carry
		if cur.Val >= 10 {
			cur.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	l := l1
	if l1 == nil {
		l = l2
	}

	cur.Next = l
	for carry != 0 && l != nil {
		cur = l
		l.Val += carry
		if l.Val >= 10 {
			l.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		l = l.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head.Next

}
