package p1256

func encode(num int) string {
	// f(0) = "", f(1) = 0, f(2) = 1, f(3) = 00, f(4) = 01, f(5) = 10, f(6) = 11, f(7) = 000
	// 1 ~ 2 = 1 bit, 3 ~ 6 = 2 bit, 7 ~ 15 = 3 bit
	// 00000 -> 11111
	// "" 0 1 10 11 100 101 110 111
	if num == 0 {
		return ""
	}

	h := uint(1)

	var pos int

	for pos+(1<<h) < num {
		pos += 1 << h
		h++
	}

	// pos + (1 << h) >= num
	H := int(h)
	buf := make([]byte, H)

	// 00000 -> 11111

	var dfs func(i int, n int)

	dfs = func(i int, n int) {
		if i > 0 {
			x := 1 << uint(i-1)
			if n > x {
				buf[i-1] = '1'
				dfs(i-1, n-x)
			} else {
				buf[i-1] = '0'
				dfs(i-1, n)
			}
		}
	}

	dfs(H, num-pos)

	for i, j := 0, H-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}
