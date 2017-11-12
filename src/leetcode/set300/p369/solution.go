package main

/**
Given a non-negative number represented as a singly linked list of digits, plus one to the number.

The digits are stored such that the most significant digit is at the head of the list.

Example:
Input:
1->2->3

Output:
1->2->4
*/

//ListNode is a holder
type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	var add func(*ListNode) (*ListNode, int)
	add = func(node *ListNode) (*ListNode, int) {
		if node == nil {
			return nil, 1
		}
		next, carray := add(node.Next)
		sum := node.Val + carray
		if sum >= 10 {
			return &ListNode{sum - 10, next}, 1
		}

		return &ListNode{sum, next}, 0
	}

	next, carray := add(head)
	if carray == 1 {
		return &ListNode{1, next}
	}
	return next
}
