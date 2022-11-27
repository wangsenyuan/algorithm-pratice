package p2847

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var loop func(node *ListNode) (int, *ListNode)

	loop = func(node *ListNode) (int, *ListNode) {
		if node.Next == nil {
			return node.Val, node
		}
		mx, tmp := loop(node.Next)
		if node.Val < mx {
			return mx, tmp
		}
		node.Next = tmp
		return node.Val, node
	}

	_, res := loop(head)

	return res
}
