package p328

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	nodeX := head.Next
	odd, even := head, head.Next

	tmp := even
	for tmp.Next != nil {
		a := tmp.Next
		odd.Next = a
		odd = a
		tmp = odd
		if tmp.Next != nil {
			b := tmp.Next
			even.Next = b
			even = b
			tmp = even
		}
	}
	odd.Next = nodeX
	even.Next = nil
	return head
}
