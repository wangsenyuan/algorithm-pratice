package p2058

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {

	if head == nil || head.Next == nil {
		return []int{-1, -1}
	}
	first := -1
	minDist := -1
	maxDist := -1
	prev := head
	node := head.Next

	isLocal := func(a, b, c *ListNode) bool {
		if a.Val < b.Val && b.Val > c.Val {
			return true
		}
		if a.Val > b.Val && b.Val < c.Val {
			return true
		}
		return false
	}
	prevLocal := -1
	pos := 1
	for node.Next != nil {
		if isLocal(prev, node, node.Next) {
			if prevLocal > 0 {
				if minDist < 0 || minDist > pos-prevLocal {
					minDist = pos - prevLocal
				}
			}

			if first > 0 {
				maxDist = pos - first
			} else {
				first = pos
			}

			prevLocal = pos
		}
		prev = node
		node = node.Next
		pos++
	}

	return []int{minDist, maxDist}
}
