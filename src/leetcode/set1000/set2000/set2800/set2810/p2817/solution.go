package p2817

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	var process func(node *ListNode) (*ListNode, int)
	process = func(node *ListNode) (res *ListNode, carry int) {
		if node == nil {
			return nil, 0
		}
		tmp, nc := process(node.Next)
		cur := node.Val*2 + nc
		carry = cur / 10
		res = &ListNode{cur % 10, tmp}
		return
	}
	newHead, carry := process(head)

	if carry > 0 {
		newHead = &ListNode{carry, newHead}
	}
	return newHead
}
