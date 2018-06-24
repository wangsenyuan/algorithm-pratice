package p856

func scoreOfParentheses(S string) int {
	n := len(S)
	if n == 2 {
		return 1
	}

	j := pair(S)
	if j == n-1 {
		return 2 * scoreOfParentheses(S[1:j])
	}
	A := S[0 : j+1]
	B := S[j+1:]
	return scoreOfParentheses(A) + scoreOfParentheses(B)
}

func pair(S string) int {
	var level int

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			level++
		} else {
			level--
		}
		if level == 0 {
			return i
		}
	}

	return -1
}
