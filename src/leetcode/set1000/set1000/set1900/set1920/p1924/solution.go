package p1924

func isDecomposable(s string) bool {
	var j int
	var two int
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[i-1] {
			l := i - j

			if l%3 == 1 {
				return false
			}
			if l%3 == 2 {
				two++
			}

			j = i
		}
	}

	return two == 1
}
