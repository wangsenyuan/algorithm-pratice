package p1171

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var first *ListNode
	var prev *ListNode
	cur := head

	for cur != nil {
		sum := cur.Val
		next := cur.Next
		for next != nil && sum != 0 {
			sum += next.Val
			next = next.Next
		}
		if sum == 0 {
			// need to remove
			cur = next
		} else {
			// keep and advance
			node := &ListNode{cur.Val, nil}
			if first == nil {
				first = node
				prev = node
			} else {
				prev.Next = node
				prev = prev.Next
			}
			cur = cur.Next
		}
	}

	return first
}
