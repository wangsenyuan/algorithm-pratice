package p2138

func divideString(s string, k int, fill byte) []string {
	n := len(s)
	x := n / k
	if x*k < n {
		x++
	}
	m := x * k
	buf := make([]byte, m)
	copy(buf, s)
	for i := n; i < m; i++ {
		buf[i] = fill
	}
	res := make([]string, x)
	for i := 0; i < x; i++ {
		res[i] = string(buf[i*k : (i+1)*k])
	}
	return res
}
