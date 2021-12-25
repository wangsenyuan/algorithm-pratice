package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n, q := readTwoNums(scanner)
	roads := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		roads[i] = readNNums(scanner, 3)
	}
	qs := make([][]int, q)
	for i := 0; i < q; i++ {
		qs[i] = readNNums(scanner, 2)
	}

	res := solve(n, roads, q, qs)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, roads [][]int, q int, queries [][]int) []int64 {
	conn := make([][]int, n)
	cost := make([][]int64, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, 10)
		cost[i] = make([]int64, 0, 10)
	}

	for _, road := range roads {
		u, v, c := road[0]-1, road[1]-1, int64(road[2])
		conn[u] = append(conn[u], v)
		cost[u] = append(cost[u], c)
		conn[v] = append(conn[v], u)
		cost[v] = append(cost[v], c)
	}

	open := make([]int, n)
	close := make([]int, n)
	D := 21

	P := make([][]int, D)
	for i := 0; i < D; i++ {
		P[i] = make([]int, n)
	}

	depth := make([]int, n)
	fp := make([]int64, n)
	up := make([]int64, n)
	down := make([]int64, n)

	PX := make([]int, n)

	var dfs func(p, u int, d int, price int64, time *int)

	dfs = func(p, u int, d int, price int64, time *int) {
		depth[u] = d
		fp[u] = fp[p] + price
		open[u] = *time
		*time++
		P[0][u] = p
		for i := 0; i < len(conn[u]); i++ {
			v := conn[u][i]
			c := cost[u][i]
			if p == v {
				PX[u] = i
				continue
			}
			up[v] = max(up[v], up[u]+c)
			dfs(u, v, d+1, c, time)
			if down[v]+c > down[u] {
				down[u] = down[v] + c
			}
		}

		close[u] = *time
	}

	dfs(0, 0, 0, 0, new(int))

	for d := 1; d < D; d++ {
		for i := 0; i < n; i++ {
			P[d][i] = P[d-1][P[d-1][i]]
		}
	}

	caseOne := func(a, b int) int64 {
		// a is ancestor of b
		return fp[b] - fp[a] + 2*max(up[a], down[b])
	}

	lca := func(u, v int) int {
		// u is higher than v
		for d := D - 1; d >= 0; d-- {
			if depth[v]-(1<<uint(d)) >= depth[u] {
				v = P[d][v]
			}
		}

		if u == v {
			return u
		}
		for d := D - 1; d >= 0; d-- {
			if P[d][u] != P[d][v] {
				u = P[d][u]
				v = P[d][v]
			}
		}
		return P[0][v]
	}

	caseTwo := func(a, b int) int64 {
		p := lca(a, b)

		return fp[a] + fp[b] - 2*fp[p] + 2*up[p]
	}

	ans := make([]int64, q)

	for i := 0; i < q; i++ {
		a, b := queries[i][0]-1, queries[i][1]-1
		if depth[a] > depth[b] {
			a, b = b, a
		}

		// a is higher than b

		if open[a] <= open[b] && close[b] <= close[a] {
			// a is parent of b
			ans[i] = caseOne(a, b)
		} else {
			ans[i] = caseTwo(a, b)
		}

	}

	return ans
}

func max3(a, b, c int64) int64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
