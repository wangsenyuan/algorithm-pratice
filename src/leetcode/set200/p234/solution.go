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
	slow, fast := head, head
	var rev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		rev, slow.Next, slow = slow, rev, slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}
	head = slow

	res := true

	for rev != nil {
		res = res && rev.Val == slow.Val
		tmp := rev.Next
		rev.Next, head, slow, rev = head, rev, slow.Next, tmp
	}

	return res
}
