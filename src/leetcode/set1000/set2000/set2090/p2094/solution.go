package p2094

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	a := head
	b := head

	for b != nil {
		b = b.Next
		if b.Next == nil {
			// 2 * n
			a.Next = a.Next.Next
			break
		}
		b = b.Next
		if b.Next == nil {
			// 2 * n + 1
			a.Next = a.Next.Next
			break
		}
		a = a.Next
	}

	return head
}
