package p1720

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	n := size(head)
	kk := n - (k - 1)

	a := nth(head, k)
	b := nth(head, kk)

	a.Val, b.Val = b.Val, a.Val

	return head
}

func size(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + size(head.Next)
}

func nth(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	return nth(head.Next, k-1)
}

func values(head *ListNode) []int {
	n := size(head)

	res := make([]int, n)

	tmp := head
	var i int
	for tmp != nil {
		res[i] = tmp.Val
		i++
		tmp = tmp.Next
	}

	return res
}
