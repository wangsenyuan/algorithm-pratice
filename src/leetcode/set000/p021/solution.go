package p021

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, cur *ListNode

	a, b := l1, l2

	for a != nil && b != nil {
		if a.Val <= b.Val {
			if head == nil {
				head = a
			}
			if cur != nil {
				cur.Next = a
			}
			cur = a
			a = a.Next
		} else {
			if head == nil {
				head = b
			}
			if cur != nil {
				cur.Next = b
			}
			cur = b
			b = b.Next
		}
	}

	if a != nil {
		cur.Next = a
	} else if b != nil {
		cur.Next = b
	}

	return head
}
