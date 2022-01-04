package p2112

func checkString(s string) bool {
	var i int

	for i < len(s) && s[i] == 'a' {
		i++
	}

	if i == len(s) {
		return true
	}

	for i < len(s) && s[i] == 'b' {
		i++
	}

	return i == len(s)
}
