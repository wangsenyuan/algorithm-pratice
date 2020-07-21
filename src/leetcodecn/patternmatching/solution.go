package patternmatching

func patternMatching(pattern string, value string) bool {
	if len(pattern) == 0 {
		return len(value) == 0
	}

	if len(pattern) == 1 {
		return true
	}

	var ax, bx int

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == pattern[0] {
			ax++
		} else {
			bx++
		}
	}

	//first find a & b
	n := len(value)

	check2 := func(a byte, A string, b byte, B string, x, y int) bool {
		if ax*len(A)+bx*len(B) != n {
			return false
		}
		for x < len(pattern) {
			var S string
			if pattern[x] == a {
				S = A
			} else {
				S = B
			}

			var j int

			for j < len(S) && y < len(value) && value[y] == S[j] {
				y++
				j++
			}
			if j < len(S) {
				return false
			}
			x++
		}

		return y == len(value) && x == len(pattern)
	}

	check := func(a byte, A string) bool {
		var i int

		var pos int

		for i < len(pattern) && pattern[i] == a {
			var j int

			for j < len(A) && pos < len(value) && A[j] == value[pos] {
				j++
				pos++
			}
			if j < len(A) {
				return false
			}

			i++
		}
		if i == len(pattern) {
			return pos == len(value)
		}

		b := pattern[i]
		for j := pos; j <= len(value); j++ {
			B := value[pos:j]
			if A != B && check2(a, A, b, B, i+1, j) {
				return true
			}
		}

		return false
	}

	a := pattern[0]

	for i := 0; i <= n; i++ {
		A := value[:i]

		if check(a, A) {
			return true
		}
	}

	return false
}
