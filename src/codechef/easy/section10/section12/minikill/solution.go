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

	N := readNum(scanner)

	edges := make([][]int, N-1)

	for i := 0; i < N-1; i++ {
		edges[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(N, edges))
}

func solve(N int, edges [][]int) int {
	dp := make([]int, N)

	outs := make([][]int, N)

	for i := 0; i < N; i++ {
		outs[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
		x := -1

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			if x == -1 || dp[v] > dp[x] {
				x = v
			}
		}

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			if x != v {
				dp[u] += max(dp[v], 1)
			} else {
				dp[u] += dp[v]
			}
		}
	}
	dfs(-1, 0)

	return max(dp[0], 1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(N int, edges [][]int) int {

	if N <= 2 {
		return 1
	}

	degree := make([]int, N)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		degree[u]++
		degree[v]++
	}

	outs := make([][]int, N)

	for i := 0; i < N; i++ {
		outs[i] = make([]int, 0, degree[i])
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	var dfs func(p, u int) int
	var total int
	var pp int
	dfs = func(p, u int) int {
		leaf := true
		r := 1
		for _, v := range outs[u] {
			if p == v {
				continue
			}
			leaf = false
			x := dfs(u, v)
			if x > 1 {
				r = 2
			}
		}

		if leaf {
			total++
		} else if r == 1 {
			if u == 0 && degree[u] > 1 {
				pp++
			} else if degree[u] > 2 {
				pp++
			}
		}

		if r > 1 {
			return r
		}

		return degree[u] - 1
	}

	dfs(-1, 0)

	return total - pp
}
