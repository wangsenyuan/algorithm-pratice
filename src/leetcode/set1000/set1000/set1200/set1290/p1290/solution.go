package p1290

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var res int

	cur := head

	for cur != nil {
		res = res<<1 | cur.Val
		cur = cur.Next
	}

	return res
}
