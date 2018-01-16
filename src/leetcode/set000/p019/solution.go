package p019

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	m := n + 1
	var prev *ListNode
	wnd := make([]*ListNode, n+1)
	h, t := 0, 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		wnd[t] = tmp
		t = (t + 1) % m
		if t == h {
			prev = wnd[h]
			h = (h + 1) % m
		}
	}
	if wnd[h] == head {
		head = wnd[h].Next
		wnd[h].Next = nil
		return head
	}
	prev.Next = wnd[h].Next
	return head
}
