package p1716

func maximumGain(s string, x int, y int) int {
	n := len(s)
	// 如果有   baba  => 2 * y or x + y
	// bababa => y + 2 * x
	// ababab => x + 2 * y
	var res int

	for i := 0; i < n; i++ {
		if s[i] == 'a' || s[i] == 'b' {
			var cnta, cntb int
			for i < n {
				if s[i] == 'a' {
					cnta++
					if y >= x && cntb > 0 {
						// better to have ba
						res += y
						cnta--
						cntb--
					}
				} else if s[i] == 'b' {
					cntb++
					if x > y && cnta > 0 {
						res += x
						cnta--
						cntb--
					}
				} else {
					break
				}
				i++
			}
			if cnta == 0 || cntb == 0 {
				// no contribution
				continue
			}
			// y > x but pattern is aaaabbb
			// or x > y but pattern is bbbaaa
			//  aaababb
			res += min(x, y) * min(cnta, cntb)
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
