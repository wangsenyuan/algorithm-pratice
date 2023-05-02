package p2663

func smallestBeautifulString(s string, k int) string {
	// 只需要检查length 2 & 3的字符串即可
	n := len(s)

	get := func(x byte) int {
		return int(x - 'a')
	}

	check := func(s []byte, p int, want int) bool {
		// can put want at p position
		if p > 0 && get(s[p-1]) == want {
			return false
		}
		if p > 1 && get(s[p-2]) == want {
			return false
		}
		return true
	}
	// abcab
	answer := func(buf []byte, p int, want int) string {
		buf[p] = byte('a' + want)
		for i := p + 1; i < n; i++ {
			// first try to set 'a', then 'b', then 'c', only three choices
			ok := false
			for x := 0; x < 3 && x < k && !ok; x++ {
				ok = check(buf, i, x)
				if ok {
					buf[i] = byte('a' + x)
				}
			}

			if !ok {
				return ""
			}
		}
		return string(buf)
	}

	ss := []byte(s)

	for i := n - 1; i >= 0; i-- {
		x := get(s[i])
		ok := false
		j := 1
		for ; x+j < k && !ok; j++ {
			ok = check(ss, i, x+j)
		}
		if !ok {
			continue
		}
		return answer(ss, i, x+j-1)
	}
	return ""
}
