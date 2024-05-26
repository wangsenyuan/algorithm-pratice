package p3159

func queryResults(limit int, queries [][]int) []int {
	occ := make(map[int]int)

	color_freq := make(map[int]int)

	var res int

	remove := func(c int) {
		if color_freq[c] == 1 {
			delete(color_freq, c)
			res--
		} else {
			color_freq[c]--
		}
	}

	add := func(c int) {
		if color_freq[c] == 0 {
			res++
		}
		color_freq[c]++
	}

	play := func(q []int) int {
		x, y := q[0], q[1]

		if v, ok := occ[x]; ok {
			// 如果之前ball x已经有颜色v
			if v == y {
				// no change
				return res
			}
			remove(v)
			add(y)
		} else {
			add(y)
		}
		occ[x] = y

		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = play(cur)
	}

	return ans
}
