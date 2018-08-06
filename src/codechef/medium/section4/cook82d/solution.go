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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N, M := readTwoNums(scanner)
	edges := make([][]int, M)
	for i := 0; i < M; i++ {
		edges[i] = readNNums(scanner, 2)
	}
	Q := readNum(scanner)
	queries := make([][]int, Q)
	for i := 0; i < Q; i++ {
		queries[i] = readNNums(scanner, 2)
	}
	res := solve(N, M, edges, Q, queries)
	for _, ans := range res {
		if ans {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

const MOD = 1e9 + 7

const MAX_N = 1e5 + 10

var powers []int64

func init() {
	powers = make([]int64, MAX_N)
	powers[0] = 1
	for i := 1; i < MAX_N; i++ {
		powers[i] = (2 * powers[i-1]) % MOD
	}
}

func add(x *int64, y int64) {
	*x += y
	if (*x) >= MOD {
		*x -= MOD
	}
}

func sub(x *int64, y int64) {
	*x -= y
	if (*x) < 0 {
		*x += MOD
	}
}

func solve(N int, M int, edges [][]int, Q int, queries [][]int) []bool {
	pref := make([]int64, M+1)
	deg := make([]int, N+1)
	var cur int64
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		if deg[u]%2 == 0 {
			add(&cur, powers[u])
		} else {
			sub(&cur, powers[u])
		}

		if deg[v]%2 == 0 {
			add(&cur, powers[v])
		} else {
			sub(&cur, powers[v])
		}
		pref[i+1] = cur
		deg[u]++
		deg[v]++
	}

	res := make([]bool, Q)

	for i, query := range queries {
		l, r := query[0], query[1]
		res[i] = pref[r] == pref[l-1]
	}

	return res
}
