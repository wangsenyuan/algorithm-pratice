package p1583

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {

	pp := make([]int, n)

	for _, p := range pairs {
		u, v := p[0], p[1]
		pp[u] = v
		pp[v] = u
	}
	var res int

	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, n)
		for j := 0; j < n-1; j++ {
			pos[i][preferences[i][j]] = j
		}
	}

	for u := 0; u < n; u++ {
		v := pp[u]
		// let's find the preference of
		var happy = true
		var i int
		for i < len(preferences[u]) && preferences[u][i] != v && happy {
			x := preferences[u][i]
			// u prefer x instead of v
			j := pos[x][u]
			k := pos[x][pp[x]]
			if j < k {
				happy = false
			}
			i++
		}

		if !happy {
			res++
		}
	}

	return res
}
