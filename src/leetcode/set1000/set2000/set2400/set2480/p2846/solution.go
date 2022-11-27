package p2846

func appendCharacters(s string, t string) int {
	var i, j int

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}

	return len(t) - j
}
