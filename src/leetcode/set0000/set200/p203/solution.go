package p203

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{-1, head}
	p := newHead

	for p.Next != nil {
		n := p.Next
		if n.Val != val {
			p = n
			continue
		}
		p.Next = n.Next
	}

	return newHead.Next
}
