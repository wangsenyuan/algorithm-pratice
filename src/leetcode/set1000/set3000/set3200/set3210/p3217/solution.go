package p3217

func getSmallestString(s string) string {
	buf := []byte(s)
	for i := 0; i+1 < len(s); i++ {
		x := int(buf[i] - '0')
		y := int(buf[i+1] - '0')
		if x&1 == y&1 && x > y {
			buf[i], buf[i+1] = buf[i+1], buf[i]
			break
		}
	}
	return string(buf)
}
