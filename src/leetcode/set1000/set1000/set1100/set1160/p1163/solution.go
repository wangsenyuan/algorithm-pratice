package p1163

func lastSubstring(s string) string {
	n := len(s)
	i := 0
	j := 1
	offset := 0
	for i+offset < n && j+offset < n {
		if s[i+offset] == s[j+offset] {
			offset++
		} else {
			if s[i+offset] < s[j+offset] {
				i += offset + 1
			} else {
				j += offset + 1
			}
			if i == j {
				i++
			}
			offset = 0
		}
	}
	return s[min(i, j):]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lastSubstring1(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var xi int
	for i := 1; i < n; i++ {
		if s[i] > s[xi] {
			xi = i
		}
	}

	var cand []int

	for i := 0; i < n; i++ {
		if s[i] == s[xi] {
			cand = append(cand, i)
		}
	}
	l := 1
	for len(cand) > 1 {

		var j int
		for i := 0; i < len(cand); i++ {
			if cand[i]+l < n {
				cand[j] = cand[i]
				j++
			}
		}
		cand = cand[:j]

		c := s[cand[0]+l]

		for i := 1; i < len(cand); i++ {
			if s[cand[i]+l] > c {
				c = s[cand[i]+l]
			}
		}

		j = 0
		for i := 0; i < len(cand); i++ {
			if s[cand[i]+l] == c {
				cand[j] = cand[i]
				j++
			}
		}

		cand = cand[:j]

		j = 0
		for i := 1; i < len(cand); i++ {
			if cand[j]+l < cand[i] {
				j++
				cand[j] = cand[i]
			}
		}
		cand = cand[:j+1]
		l++
	}

	return s[cand[0]:]
}
