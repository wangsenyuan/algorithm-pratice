package p1526

func restoreString(s string, indices []int) string {
	buf := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		buf[indices[i]] = s[i]
	}

	return string(buf)
}
