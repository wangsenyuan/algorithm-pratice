package p142

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for slow != nil && fast != nil {
		if slow.Next == nil {
			return nil
		}
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next

		if slow == fast {
			fast = head

			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
