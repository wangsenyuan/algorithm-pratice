package p125

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	i, j := 0, len(s)-1

	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if !equalsIgnoreCase(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func equalsIgnoreCase(a, b byte) bool {
	if a == b {
		return true
	}

	if toLower(a) == toLower(b) {
		return true
	}
	return false
}

func toLower(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a - 'A' + 'a'
	}
	return a
}
