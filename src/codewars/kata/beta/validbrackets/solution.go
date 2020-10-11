package validbrackets

func ValidMixedBracketsParentheses(s string) bool {
	// Your code goes here
	n := len(s)
	st := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			st[p] = i
			p++
			continue
		}
		if p == 0 {
			return false
		}
		j := st[p-1]
		if s[i] == '}' && s[j] != '{' {
			return false
		}
		if s[i] == ')' && s[j] != '(' {
			return false
		}
		if s[i] == ']' && s[j] != '[' {
			return false
		}
		p--
	}

	return p == 0
}
