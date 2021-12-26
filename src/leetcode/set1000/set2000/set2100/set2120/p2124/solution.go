package p2124

func executeInstructions(n int, startPos []int, s string) []int {
	m := len(s)
	ans := make([]int, m)

	check := func(i int) int {
		x, y := startPos[0], startPos[1]
		var cnt int
		for i < m {
			if s[i] == 'L' {
				y--
			} else if s[i] == 'R' {
				y++
			} else if s[i] == 'U' {
				x--
			} else {
				x++
			}
			if x < 0 || x == n || y < 0 || y == n {
				break
			}
			cnt++
			i++
		}

		return cnt
	}

	for i := 0; i < m; i++ {
		ans[i] = check(i)
	}

	return ans
}
