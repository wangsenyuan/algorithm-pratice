package p1247

func minimumSwap(s1 string, s2 string) int {
	var x, y int
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'x' {
			x++
		} else {
			y++
		}
	}

	for i := 0; i < len(s2); i++ {
		if s2[i] == 'x' {
			x++
		} else {
			y++
		}
	}

	if x%2 != 0 {
		// can't distribute x or y evenly
		return -1
	}
	// xy & yx  => 2
	// xx & yy  => 1

	var a, b int

	for i := 0; i < len(s1); i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			a++
		} else if s1[i] == 'y' && s2[i] == 'x' {
			b++
		}
	}

	// xxxxxx vs yyyyy =>
	res := a/2 + b/2
	a %= 2
	b %= 2

	if a != b {
		return -1
	}

	if a == 1 {
		res += 2
	}

	return res
}
