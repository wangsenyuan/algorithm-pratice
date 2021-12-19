package p2108

func firstPalindrome(words []string) string {

	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
