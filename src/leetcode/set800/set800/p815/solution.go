package p815

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	rr := make(map[int][]int)

	n := len(routes)

	for i, route := range routes {

		for _, stop := range route {

			if _, found := rr[stop]; !found {
				rr[stop] = make([]int, 0, n)
			}
			rr[stop] = append(rr[stop], i)
		}
	}

	var travel func(stop int, route int) int

	cache := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		cache[i] = make(map[int]int)
	}
	visited := make(map[int]bool)
	travel = func(stop int, route int) int {
		visited[route] = true
		if _, found := cache[route][stop]; !found {
			min := 1 << 30
			for _, w := range routes[route] {

				if w == T {
					min = 1
					break
				}
				nextRoutes := rr[w]
				for _, x := range nextRoutes {
					if visited[x] {
						continue
					}
					tmp := travel(w, x) + 1
					if tmp < min {
						min = tmp
					}
				}
			}

			cache[route][stop] = min
		}
		visited[route] = false
		return cache[route][stop]
	}
	res := 1 << 30
	first := rr[S]
	for _, route := range first {
		tmp := travel(S, route)
		if tmp < res {
			res = tmp
		}
	}

	if res < 1<<30 {
		return res
	}
	return -1
}

func numBusesToDestination1(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	n := len(routes)
	has := make([]map[int]bool, n)

	for i, route := range routes {
		has[i] = make(map[int]bool)
		for _, stop := range route {
			has[i][stop] = true
		}
	}

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = -1
		}
		f[i][i] = 0
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for _, stop := range routes[i] {
				if has[j][stop] {
					f[i][j] = 1
					f[j][i] = 1
					break
				}
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if f[i][k] >= 0 && f[k][j] >= 0 {
					if f[i][j] == -1 || f[i][j] > f[i][k]+f[k][j] {
						f[i][j] = f[i][k] + f[k][j]
					}
				}
			}
		}
	}

	ans := -1

	for i := 0; i < n; i++ {
		if has[i][S] {
			for j := 0; j < n; j++ {
				if has[j][T] && f[i][j] >= 0 {
					if ans == -1 || ans > f[i][j]+1 {
						ans = f[i][j] + 1
					}
				}
			}
		}
	}

	return ans
}
