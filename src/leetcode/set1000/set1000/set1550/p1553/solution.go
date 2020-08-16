package p1553

func minDays(n int) int {
	mem := make(map[int]int)

	var loop func(x int) int

	loop = func(x int) int {
		if x == 0 {
			return 0
		}

		if x == 1 {
			return 1
		}

		if v, found := mem[x]; found {
			return v
		}

		res := min(loop(x/2)+x%2+1, loop(x/3)+x%3+1)

		mem[x] = res

		return res
	}

	return loop(n)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
