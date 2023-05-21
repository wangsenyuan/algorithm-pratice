package p2698

func makeSmallestPalindrome(s string) string {
	buf := []byte(s)

	n := len(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if buf[i] == buf[j] {
			continue
		}
		if buf[i] > buf[j] {
			buf[i] = buf[j]
		} else {
			buf[j] = buf[i]
		}
	}
	return string(buf)
}
