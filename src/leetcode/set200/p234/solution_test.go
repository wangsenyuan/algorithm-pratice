package main

import "testing"

func TestIsPalindrome1(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	if isPalindrome(head) {
		t.Errorf("1 -> 2 is not palindrome, but give true")
	}
}
