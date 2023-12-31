package p2981

func maximumLength(s string) int {
	records := make([]map[int]int, 26)

	var res int = -1

	add := func(x int, y int, v int) {
		if records[x] == nil {
			records[x] = make(map[int]int)
		}
		records[x][y] += v
		if records[x][y] >= 3 {
			res = max(res, y)
		}
	}

	for i := 0; i < len(s); {
		j := i
		for i < len(s) && s[i] == s[j] {
			i++
		}
		x := int(s[j] - 'a')
		add(x, i-j, 1)
		if i-j >= 2 {
			add(x, i-j-1, 2)
		}

		if i-j >= 3 {
			add(x, i-j-2, 3)
		}
	}

	return res
}
