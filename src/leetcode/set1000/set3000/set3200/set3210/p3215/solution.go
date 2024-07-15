package p3215

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	vis := make(map[int]int)
	for _, num := range nums {
		vis[num]++
	}

	newHead := new(ListNode)
	prev := newHead
	for node := head; node != nil; node = node.Next {
		if vis[node.Val] > 0 {
			continue
		}
		prev.Next = node
		prev = node
	}

	prev.Next = nil

	return newHead.Next
}
