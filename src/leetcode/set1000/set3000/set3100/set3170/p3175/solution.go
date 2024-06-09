package p3175

func clearDigits(s string) string {
	n := len(s)
	stack := make([]byte, n)
	var top int
	for i := 0; i < n; i++ {
		if isNum(s[i]) {
			if top > 0 {
				top--
			}
		} else {
			stack[top] = s[i]
			top++
		}
	}

	return string(stack[:top])
}

func isNum(x byte) bool {
	return x >= '0' && x <= '9'
}
