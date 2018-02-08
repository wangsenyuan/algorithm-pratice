package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for x := 0; x < t; x++ {
		n, m := readTwoNums(scanner)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 3)
		}
		s, t := readTwoNums(scanner)
		res := solve(n, m, edges, s, t)
		if res == -1 {
			fmt.Println("-1")
		} else {
			fmt.Printf("%.7f\n", res)
		}
	}
}

const INF = 1e18

func solve(n int, m int, edges [][]int, s, t int) float64 {
	s--
	t--
	graph := make([][]Edge, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]Edge, 0, 10)
	}

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		graph[u] = append(graph[u], Edge{u, v, w})
	}

	conn := make([][]bool, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]bool, n)
	}

	var dfs func(v int, p int, flag int)

	visited := make([]int, n)

	dfs = func(v int, p int, flag int) {
		visited[v] = flag
		conn[p][v] = true

		for _, edge := range graph[v] {
			w := edge.v
			if visited[w] != flag {
				dfs(w, p, flag)
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(i, i, i+1)
	}

	if !conn[s][t] {
		return -1
	}

	d := make([]float64, n)

	check := func(avg float64) bool {
		for i := 0; i < n; i++ {
			d[i] = INF
		}
		d[s] = 0
		for k := 1; k < n; k++ {
			for _, edge := range edges {
				u, v, w := edge[0]-1, edge[1]-1, float64(edge[2])
				w -= avg
				if d[u]+w < d[v] {
					d[v] = d[u] + w
				}
			}
		}
		if d[t] <= 0 {
			return true
		}

		for _, edge := range edges {
			u, v, w := edge[0]-1, edge[1]-1, float64(edge[2])
			w -= avg
			if d[u]+w < d[v] && conn[s][v] && conn[v][t] {
				return true
			}
		}
		return false
	}

	const EPS = 1e-9

	left := -100.0
	right := 100.0

	for right-left > EPS {
		mid := (left + right) / 2.0
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}

type Edge struct {
	u int
	v int
	w int
}
