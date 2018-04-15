package p817

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	if head == nil {
		return 0
	}
	gm := make(map[int]bool)

	for _, g := range G {
		gm[g] = true
	}

	var ans int
	for node := head; node != nil; node = node.Next {
		if node.Next == nil {
			if gm[node.Val] {
				ans++
			}
		} else {
			if gm[node.Val] && !gm[node.Next.Val] {
				ans++
			}
		}
	}

	return ans
}
