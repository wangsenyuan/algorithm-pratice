package p2381

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	cnt := make([]int, n+1)

	for _, shift := range shifts {
		s, e, d := shift[0], shift[1], shift[2]
		if d == 0 {
			cnt[s] += 25
			cnt[e+1] -= 25
		} else {
			cnt[s]++
			cnt[e+1]--
		}
	}
	buf := []byte(s)
	for i := 0; i < n; i++ {
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
		tmp := (cnt[i]%26 + 26) % 26
		x := int(buf[i] - 'a')
		x = (x + tmp) % 26
		buf[i] = byte('a' + x)
	}

	return string(buf)
}
