package p382

import "math/rand"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	size int
	head *ListNode
}

/** @param head The linked list's head. Note that the head is guanranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	rand.Seed(41)
	return Solution{sizeOf(head), head}
}

func sizeOf(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + sizeOf(head.Next)
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	x := rand.Intn(this.size)
	tmp := this.head
	for i := 0; i < x; i++ {
		tmp = tmp.Next
	}
	return tmp.Val
}
