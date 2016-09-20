package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Printf("%s is palindrome? %v", head, isPalindrome(head))
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	if head == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> %v", head.Val, head.Next.String())
}

func isPalindrome(head *ListNode) bool {
	sz := size(head)

	if sz <= 1 {
		return true
	}

	half := sz / 2

	a := head
	for i := 0; i < half-1; i++ {
		a = a.Next
	}

	b, _ := reverse(a.Next)

	for c, d := head, b; 0 < half; half-- {
		if c.Val != d.Val {
			break
		}
		c, d = c.Next, d.Next
	}

	b, _ = reverse(b)

	a.Next = b

	return half == 0
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	newHead, last := reverse(head.Next)
	last.Next = head
	head.Next = nil
	return newHead, head
}

func size(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + size(head.Next)
}
