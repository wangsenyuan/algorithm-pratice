package p1545

func longestAwesome(s string) int {
	pos := make([]int, 1<<11)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0

	var res int
	var cur int

	for i := 1; i <= len(s); i++ {
		x := int(s[i-1] - '0')

		cur ^= (1 << x)

		if pos[cur] >= 0 {
			res = max(res, i-pos[cur])
		}

		for j := 0; j < 10; j++ {
			// if j is odd
			tmp := cur ^ (1 << j)
			if pos[tmp] >= 0 {
				res = max(res, i-pos[tmp])
			}
		}

		if pos[cur] < 0 {
			pos[cur] = i
		}
	}

	return res
}
