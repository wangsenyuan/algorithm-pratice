package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)

		f := make([]int, 0, n)
		g := make(map[int][]int)

		for i := 0; i < n; i++ {
			var x = 0
			fmt.Scanf("%d", &x)
			x = x - 1
			f = append(f, x)
			if edges, ok := g[x]; !ok {
				edges = make([]int, 0, n)
				g[x] = append(edges, i)
			} else {
				g[x] = append(edges, i)
			}
		}
		r := play(f, n, g)
		fmt.Printf("Case #%d: %d\n", i, r)
	}
}

func play(f []int, n int, g map[int][]int) int {

	var loop func(int, []int) int

	loop = func(v int, path []int) int {
		j := sliceIndex(len(path), func(i int) bool {
			return path[i] == v
		})
		if j >= 0 {
			return len(path) - j
		}
		return loop(f[v], append(path, v))
	}

	var r1 = 0
	for i := 0; i < n; i++ {
		path := make([]int, 0, n)
		path = append(path, i)
		r1 = max(r1, loop(f[i], path))
	}

	var dfs func(int, int, int) int

	dfs = func(p, i, d int) int {
		xs := g[i]
		r := d
		for _, x := range xs {
			if p == x {
				continue
			}
			r = max(r, dfs(i, x, d+1))
		}
		return r
	}

	var r2 = 0
	for i := 0; i < n; i++ {
		if f[f[i]] != i {
			continue
		}

		r2 += dfs(i, f[i], 0) + 1
	}

	return max(r1, r2)
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
