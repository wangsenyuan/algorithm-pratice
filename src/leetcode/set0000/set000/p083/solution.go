package p083

/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			cur.Next = next.Next
			next = next.Next
		}
		cur = cur.Next
	}

	return head
}
