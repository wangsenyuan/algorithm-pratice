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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readNum(scanner)

	K := readNNums(scanner, n)

	edges := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, edges, K))
}

func solve(n int, edges [][]int, K []int) int {

	degree := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		degree[u-1]++
		degree[v-1]++
	}

	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, degree[i])
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}
	MAX_K := 1024

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, MAX_K)
		for j := 0; j < MAX_K; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(p, u, x int) int

	var dfsSkip func(p, u, x, j int) int

	dfs = func(p, u, x int) int {
		if dp[u][x] >= 0 {
			return dp[u][x]
		}

		var s1, s2 int

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			s1 += dfs(u, v, x)

			if x^K[u] == 0 {
				s2 += dfs(u, v, 0)
			} else {
				s2 += dfsSkip(u, v, x^K[u], x^K[u])
			}
		}
		s2 += K[u]

		dp[u][x] = max(s1, s2)

		return dp[u][x]
	}

	dfsSkip = func(p, u, x, j int) int {
		var s int

		if j == 1 {
			for _, v := range outs[u] {
				if p != v {
					s += dfs(u, v, x)
				}
			}
		} else {
			for _, v := range outs[u] {
				if p != v {
					s += dfsSkip(u, v, x, j-1)
				}
			}
		}
		return s
	}

	return dfs(-1, 0, 0)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
