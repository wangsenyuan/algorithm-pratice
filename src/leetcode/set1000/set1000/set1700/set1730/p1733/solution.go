package p1733

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	known := make([]map[int]bool, len(languages))

	for i, cur := range languages {
		known[i] = make(map[int]bool)
		for _, j := range cur {
			known[i][j] = true
		}
	}

	left := make([]int, 0, len(friendships))

	for i, cur := range friendships {
		u, v := cur[0], cur[1]
		u--
		v--
		var ok bool
		for j := 1; j <= n; j++ {
			if known[u][j] && known[v][j] {
				ok = true
				break
			}
		}
		if !ok {
			left = append(left, i)
		}
	}

	if len(left) == 0 {
		return 0
	}

	vis := make([]int, len(languages))

	check := func(l int) int {
		var cnt int

		for _, i := range left {
			u, v := friendships[i][0], friendships[i][1]
			u--
			v--
			if vis[u] != l {
				vis[u] = l
				if !known[u][l] {
					cnt++
				}
			}
			if vis[v] != l {
				vis[v] = l
				if !known[v][l] {
					cnt++
				}
			}
		}

		return cnt
	}

	best := 500

	for i := 1; i <= n; i++ {
		tmp := check(i)
		if tmp < best {
			best = tmp
		}
	}
	return best
}
