package problem3

func robot(command string, obstacles [][]int, x int, y int) bool {
	n := len(command)
	var up int
	var rt int

	dp := make([]int, n+1)
	fp := make([]int, n+1)
	for i := 0; i < n; i++ {
		if command[i] == 'U' {
			up++
		} else {
			rt++
		}
		dp[i+1] = up
		fp[i+1] = rt
	}

	// y = up * m + rx
	// x = rt * m + ry
	checkInRoad := func(u int, v int) int {
		if up == 0 {
			if v != 0 {
				return -1
			}
			return v
		}
		if rt == 0 {
			if u != 0 {
				return -1
			}
			return u
		}

		m := (u + v) / n
		v -= m * up
		u -= m * rt
		i := u + v
		if i >= 0 && i < n && dp[i] == v && fp[i] == u {
			return m*n + i
		}

		return -1
	}
	steps := checkInRoad(x, y)
	if steps < 0 {
		return false
	}

	for _, obstacle := range obstacles {
		u, v := obstacle[0], obstacle[1]
		if u > x || v > y {
			continue
		}

		tmp := checkInRoad(u, v)
		if tmp >= 0 && tmp < steps {
			return false
		}
	}

	return true
}
