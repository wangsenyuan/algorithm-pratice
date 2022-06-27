package p2315

func countAsterisks(s string) int {
	var open bool

	var cnt int
	var last int
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if !open {
				cnt++
			} else {
				last++
			}
		} else if s[i] == '|' {
			open = !open
			last = 0
		}
	}

	if open {
		cnt += last
	}

	return cnt
}
