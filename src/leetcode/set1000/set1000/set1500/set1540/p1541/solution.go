package p1541

func minInsertions(s string) int {
	n := len(s)
	var cnt int
	var res int
	for i := 0; i < n; {
		if s[i] == '(' {
			cnt++
			i++
		} else {
			// s[i] == ')'
			if i+1 == n || s[i+1] == '(' {
				i++
				res++
			} else {
				i += 2
			}

			if cnt > 0 {
				cnt--
			} else {
				res++
			}
		}
	}

	return res + 2*cnt
}
