package p1703

func halvesAreAlike(s string) bool {
	n := len(s)
	var a, b int

	for i := 0; i < n; i++ {
		if !isVowl(s[i]) {
			continue
		}
		if i*2 < n {
			a++
		} else {
			b++
		}
	}

	return a == b
}

var vowls = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func isVowl(c byte) bool {
	for i := 0; i < len(vowls); i++ {
		if c == vowls[i] {
			return true
		}
	}
	return false
}
