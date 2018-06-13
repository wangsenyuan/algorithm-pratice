package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, m int

		edges := make([][]int, n-1)
		fmt.Scanf("%d %d", &n, &m)

		for i := 0; i < n-1; i++ {
			var a, b, c int
			fmt.Scanf("%d %d %d", &a, &b, &c)
			edges[i] = []int{a, b, c}
		}

		special := make([]bool, n)

		for i := 0; i < m; i++ {
			var x int
			fmt.Scanf("%d", &x)
			special[x-1] = true
		}

		a, b := solve(n, m, edges, special)

		fmt.Printf("%d %d\n", a, b)
		t--
	}
}

func solve(n int, m int, edges [][]int, special []bool) (int64, int64) {
	grid := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		grid[i] = make(map[int]int)
	}

	for _, edge := range edges {
		a, b, c := edge[0]-1, edge[1]-1, edge[2]
		grid[a][b] = c
		grid[b][a] = c
	}

	sum := make([]int64, n)
	var dfs func(p, v int)
	var ans int64
	dfs = func(p, v int) {
		if special[v] {
			sum[v] = 1
		}

		for u, d := range grid[v] {
			if u == p {
				continue
			}
			dfs(v, u)
			ans += int64(2) * sum[u] * (int64(m) - sum[u]) * int64(d)
			sum[v] += sum[u]
		}

	}

	dfs(-1, 0)

	a := ans
	b := int64(m) * int64(m)
	d := gcd(a, b)

	return a / d, b / d
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
