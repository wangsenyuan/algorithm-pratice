package p5558

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var x, y *ListNode
	cur := list1
	for i := 0; ; i++ {
		if i+1 == a {
			x = cur
		}
		if i == b {
			y = cur
			break
		}
		cur = cur.Next
	}

	x.Next = list2

	cur = list2

	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = y.Next

	return list1
}
