package p3210

func getEncryptedString(s string, k int) string {
	n := len(s)
	k %= n
	buf := []byte(s)

	for i := 0; i < n; i++ {
		j := (i + k) % n
		buf[i] = s[j]
	}

	return string(buf)
}
