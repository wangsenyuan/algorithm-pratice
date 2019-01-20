package p237

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node != nil {
		var prev *ListNode
		for node.Next != nil {
			prev = node
			node.Val = node.Next.Val
			node = node.Next
		}
		prev.Next = nil
	}
}
