package p2807

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var res *ListNode
	var prev *ListNode
	cur := head

	for cur != nil {
		node := new(ListNode)
		node.Val = cur.Val
		if prev != nil {
			prev.Next = node
		} else {
			res = node
		}
		nxt := cur.Next

		if nxt != nil {
			g := gcd(cur.Val, nxt.Val)
			node.Next = new(ListNode)
			node.Next.Val = g
			node = node.Next
		}
		prev = node
		cur = nxt
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
