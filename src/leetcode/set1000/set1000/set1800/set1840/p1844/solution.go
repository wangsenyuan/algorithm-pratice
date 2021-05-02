package p1844

func replaceDigits(s string) string {
	buf := []byte(s)

	for i := 1; i < len(s); i += 2 {
		buf[i] = shift(buf[i-1], int(s[i]-'0'))
	}
	return string(buf)
}

func shift(x byte, i int) byte {
	return byte(int(x) + i)
}
