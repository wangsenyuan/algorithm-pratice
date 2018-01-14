package p125

import "testing"

func TestSample1(t *testing.T) {
	res := isPalindrome("a.")
	if !res {
		t.Errorf("sample a. should be palindrome, but got not")
	}
}

func TestSample2(t *testing.T) {
	res := isPalindrome("aa")
	if !res {
		t.Errorf("sample a. should be palindrome, but got not")
	}
}
