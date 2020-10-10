package p141

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for slow != nil && fast != nil {
		if slow.Next == nil {
			return false
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}

		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
