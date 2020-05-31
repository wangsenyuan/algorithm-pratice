package p1461

func hasAllCodes(s string, k int) bool {
	n := len(s)
	m := 1 << k
	h := m >> 1

	mem := make(map[int]bool)
	var num int
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		num = num<<1 | x

		if i >= k-1 {
			mem[num] = true
			num &= (h - 1)
		}
	}

	return len(mem) == m
}
