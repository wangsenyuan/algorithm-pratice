package p1785

func checkOnesSegment(s string) bool {
	if len(s) == 0 {
		return false
	}

	var i int

	for i < len(s) && s[i] == '1' {
		i++
	}
	if i == len(s) {
		return true
	}
	for i < len(s) && s[i] == '0' {
		i++
	}

	return i == len(s)
}
