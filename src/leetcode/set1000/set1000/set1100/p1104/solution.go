package p1104

func pathInZigZagTree(label int) []int {
	var loop func(pos int, lmt int, odd bool)

	findNum := func(pos int, lmt int, odd bool) int {
		if odd {
			return pos
		}

		return lmt - 1 - (pos - lmt/2)
	}

	var path []int

	loop = func(pos int, lmt int, odd bool) {
		if pos > 0 {
			path = append(path, findNum(pos, lmt, odd))
			loop(pos/2, lmt/2, !odd)
		}
	}

	var i int
	lmt := 1
	for lmt <= label {
		i++
		lmt *= 2
	}
	pos := label
	if i%2 == 0 {
		pos = lmt - 1 - (label - lmt/2)
	}

	loop(pos, lmt, i%2 == 1)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
