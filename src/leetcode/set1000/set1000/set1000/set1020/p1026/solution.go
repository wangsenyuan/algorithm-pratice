package p1026

func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	res := make([]bool, n)

	for i := 0; i < n; i++ {
		res[i] = check(queries[i], pattern)
	}

	return res
}

func check(query string, pattern string) bool {
	m := len(query)
	n := len(pattern)

	var i int
	var j int

	for j < n {
		for i < m && query[i] != pattern[j] && isLower(query[i]) {
			i++
		}
		if i == m {
			return false
		}

		if query[i] != pattern[j] {
			return false
		}

		// query[i] == pattern[j]
		i++
		j++
	}

	for i < m && isLower(query[i]) {
		i++
	}

	return i == m
}

func isLower(letter byte) bool {
	return letter >= 'a' && letter <= 'z'
}
