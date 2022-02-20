package p2181

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var cur *ListNode
	var res *ListNode
	node := head

	for node != nil {
		if node.Val == 0 {
			node = node.Next
			continue
		}
		tmp := new(ListNode)
		for node != nil && node.Val != 0 {
			tmp.Val += node.Val
			node = node.Next
		}
		if cur != nil {
			cur.Next = tmp
		}
		cur = tmp
		if res == nil {
			res = cur
		}
	}

	return res
}
