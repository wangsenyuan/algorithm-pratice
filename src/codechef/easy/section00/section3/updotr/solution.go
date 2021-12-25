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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N := readNum(scanner)

		A := readNNums(scanner, N)
		P := readNNums(scanner, N-1)
		Q := readNum(scanner)
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 2)
		}
		res := solve(N, A, P, Q, queries)
		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

func solve(N int, a []int, P []int, Q int, queries [][]int) []int {
	A := make([]int, N+1)
	copy(A[1:], a)

	conn := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		conn[i] = make([]int, 0, 3)
	}

	for i := 2; i <= N; i++ {
		p := P[i-2]
		conn[p] = append(conn[p], i)
		conn[i] = append(conn[i], p)
	}

	PP := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		PP[i] = make([]int, 20)
		for j := 0; j < 20; j++ {
			PP[i][j] = -1
		}
	}

	// S[i] = the max dish index so far
	MX := make([]int, N+1)
	CT := make([]int, N+1)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		if MX[p] < A[u] {
			MX[u] = A[u]
			CT[u] = CT[p] + 1
		} else {
			MX[u] = MX[p]
			CT[u] = CT[p]
		}

		PP[u][0] = p
		for i := 1; i < 19; i++ {
			if PP[u][i-1] > 0 {
				PP[u][i] = PP[PP[u][i-1]][i-1]
			}
		}

		for _, v := range conn[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
	}
	dfs(0, 1)

	calculate := func(v, w int) int {
		if MX[v] <= w {
			return 0
		}
		if MX[0] > w {
			return CT[v]
		}

		u := v
		for i := 19; i >= 0; i-- {
			if PP[u][i] >= 0 && MX[PP[u][i]] > w {
				u = PP[u][i]
			}
		}
		return CT[v] - CT[PP[u][0]]
	}

	res := make([]int, Q)

	var last int

	for i, query := range queries {
		a, b := query[0], query[1]
		v, w := a^last, b^last

		res[i] = calculate(v, w)
		last = res[i]
	}

	return res
}
