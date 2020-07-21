package p266

import "testing"

func TestCanPermutePalindrome(t *testing.T) {
	strs := []string{"aab", "carerac"}
	for _, str := range strs {
		if !canPermutePalindrome(str) {
			t.Errorf("%v does have (at least) one palindrome permutation", str)
		}
	}
}

func TestCanNotPermutePalindrome(t *testing.T) {
	strs := []string{"code"}
	for _, str := range strs {
		if canPermutePalindrome(str) {
			t.Errorf("%v doesn't have (at least) one palindrome permutation", str)
		}
	}
}
