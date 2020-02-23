package p1358

func numberOfSubstrings(s string) int {
	n := len(s)

	pos := make([]int, 3)
	for i := 0; i < 3; i++ {
		pos[i] = -1
	}

	var res int

	for i := 0; i < n; i++ {
		// .... a...,b...,c
		x := int(s[i] - 'a')
		var a, b int
		if x == 0 {
			// a
			// we need to find last b, & c
			a, b = 1, 2
		} else if x == 1 {
			a, b = 0, 2
		} else {
			a, b = 0, 1
		}
		pos[x] = i
		j := min(pos[a], pos[b])
		res += j + 1
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
