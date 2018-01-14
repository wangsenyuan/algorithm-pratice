package p344

func reverseString(s string) string {
	if len(s) == 0 {
		return ""
	}

	bs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
