package p917

func reverseOnlyLetters(S string) string {
	n := len(S)
	bs := []byte(S)

	for i, j := 0, n-1; i < j; {
		if !isLetter(bs[i]) {
			i++
			continue
		}
		if !isLetter(bs[j]) {
			j--
			continue
		}

		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}
