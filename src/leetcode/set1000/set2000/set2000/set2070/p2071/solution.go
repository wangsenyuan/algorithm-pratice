package p2071

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := 1
	node := head

	for node != nil {
		if node.Next == nil {
			break
		}
		x := node
		var tmp int
		for tmp <= cur && x != nil {
			x = x.Next
			tmp++
		}
		if x == nil {
			tmp--
		}
		if tmp%2 == 1 {
			for node.Next != nil && tmp > 0 {
				node = node.Next
				tmp--
			}
			cur++
			continue
		}
		a := node.Next
		if a == nil || a.Next == nil {
			break
		}
		aa := a

		b := a.Next
		for b != nil && tmp > 1 {
			node.Next = b
			c := b.Next
			b.Next = a
			a = b
			b = c
			aa.Next = b
			tmp--
		}
		node = aa
		cur++
	}
	node.Next = nil

	return head
}
