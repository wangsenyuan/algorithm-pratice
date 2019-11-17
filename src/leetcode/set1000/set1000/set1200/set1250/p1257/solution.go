package p1257

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	ii := make(map[string]int)

	var index int

	for _, rgs := range regions {
		for _, rg := range rgs {
			if _, found := ii[rg]; !found {
				ii[rg] = index
				index++
			}
		}
	}

	n := len(ii)
	outs := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, rgs := range regions {
		first := rgs[0]
		a := ii[first]
		for i := 1; i < len(rgs); i++ {
			outs[a] = append(outs[a], ii[rgs[i]])
			degree[ii[rgs[i]]]++
		}
	}

	x := ii[region1]
	y := ii[region2]

	var findTarget func(u int, x int) bool
	findTarget = func(u int, x int) bool {
		if u == x {
			return true
		}

		for _, v := range outs[u] {
			if findTarget(v, x) {
				return true
			}
		}

		return false
	}

	if findTarget(x, y) {
		return region1
	}

	if findTarget(y, x) {
		return region2
	}

	var dfs func(u int, h int) int

	var best int
	var ans int
	dfs = func(u int, h int) int {
		if u == x {
			return 1
		}

		if u == y {
			return 2
		}

		var res int

		for _, v := range outs[u] {
			a := dfs(v, h+1)
			res |= a
		}

		if res == 3 && h > best {
			// it has both x & y
			ans = u
			best = h
		}
		return res
	}

	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			dfs(i, 0)
		}
	}

	for k, v := range ii {
		if v == ans {
			return k
		}
	}
	return ""
}
