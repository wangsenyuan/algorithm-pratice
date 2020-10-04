package p1545

func longestAwesome(s string) int {
	pos := make(map[int]int)
	pos[0] = 0

	var res int
	var cur int

	for i := 1; i <= len(s); i++ {
		x := int(s[i-1] - '0')

		cur ^= (1 << x)

		if j, found := pos[cur]; found {
			res = max(res, i-j)
		}

		for j := 0; j < 10; j++ {
			// if j is odd
			tmp := cur ^ (1 << j)
			if k, found := pos[tmp]; found {
				res = max(res, i-k)
			}
		}

		if _, found := pos[cur]; !found {
			pos[cur] = i
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
