package main

import (
	"bufio"
	"os"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n, k := readTwoNums(scanner)
		edges := make([][]int, n-1)
		for j := 0; j < n-1; j++ {
			edges[j] = readNNums(scanner, 2)
		}
		fmt.Printf("Case #%d: %d\n", i, solve(n, k, edges))
	}
}

const MOD = 1000000009

func solve(n int, k int, edges [][]int) int {
	conn := make([][]int, n)

	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, 3)
	}
	deg := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
		deg[u]++
		deg[v]++
	}

	var dfs func(p, u int)

	ans := int64(1)
	K := int64(k)

	dfs = func(p, u int) {
		var D int64
		if p >= 0 {
			D = int64(deg[p])
		}
		for _, v := range conn[u] {
			if p == v {
				continue
			}
			ans = (ans * (K - D)) % MOD
			D++
			dfs(u, v)
		}
	}
	//deg[0] = 0
	dfs(-1, 0)

	return int(ans)
}
