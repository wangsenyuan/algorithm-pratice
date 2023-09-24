package p2865

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	var cnt int
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		cnt += int(s[i] - '0')
		buf[i] = '0'
	}
	for i := 0; i < cnt-1; i++ {
		buf[i] = '1'
	}
	buf[n-1] = '1'
	return string(buf)
}
