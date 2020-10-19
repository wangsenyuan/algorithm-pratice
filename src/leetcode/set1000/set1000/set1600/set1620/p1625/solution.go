package p1625

func maxLengthBetweenEqualCharacters(s string) int {
	pos := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
	}

	var res int
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if pos[x] >= 0 {
			res = max(res, i-pos[x]-1)
		} else {
			pos[x] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
