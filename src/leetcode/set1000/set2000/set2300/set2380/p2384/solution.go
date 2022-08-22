package p2384

func largestPalindromic(num string) string {
	cnt := make([]int, 10)

	for i := 0; i < len(num); i++ {
		cnt[int(num[i]-'0')]++
	}
	var buf []byte

	for x := 9; x >= 0; x-- {
		if cnt[x] > 0 {
			a := cnt[x] / 2
			cnt[x] = cnt[x] % 2
			for a > 0 {
				a--
				buf = append(buf, byte(x+'0'))
			}
		}
	}
	if len(buf) > 0 && buf[0] == '0' {
		buf = buf[:0]
	}
	n := len(buf)

	rev := make([]byte, n)
	for i := 0; i < len(buf); i++ {
		rev[i] = buf[n-1-i]
	}

	for x := 9; x >= 0; x-- {
		if cnt[x] > 0 {
			buf = append(buf, byte(x+'0'))
			break
		}
	}

	buf = append(buf, rev...)

	if len(buf) == 0 {
		return "0"
	}

	return string(buf)
}
