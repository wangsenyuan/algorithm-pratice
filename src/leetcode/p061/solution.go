package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	return fmt.Sprintf("%d", ln.Val)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		n++
	}
	k = k % n

	head = rotate(head, 0, n)
	head = rotate(head, 0, k)
	head = rotate(head, k, n)

	return head
}

func rotate(head *ListNode, i, j int) *ListNode {
	if i == j {
		return head
	}
	a := head
	var prev *ListNode

	for k := 0; k < i; k++ {
		prev = a
		a = a.Next
	}

	var b *ListNode
	p := a
	c := a.Next
	for k := i; k < j; k++ {
		c = a.Next
		a.Next = b
		b = a
		a = c
	}

	p.Next = c

	if prev != nil {
		prev.Next = b
		return head
	} else {
		return b
	}
}

func main() {
	head := fromArray([]int{1, 2, 3, 4, 5, 6, 7})

	for i := 1; i <= 7; i++ {
		head = rotateRight(head, i)

		for tmp := head; tmp != nil; tmp = tmp.Next {
			fmt.Printf("%d ->", tmp.Val)
		}
		fmt.Println()
	}
}

func fromArray(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{Val: nums[0], Next: fromArray(nums[1:])}
}
