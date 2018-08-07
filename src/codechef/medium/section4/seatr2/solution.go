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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		res := solve(n, m, edges)
		fmt.Println(res)
	}
}

const MOD = 1e9 + 7

func solve(N, M int, edges [][]int) int {
	F := make([][]int64, N+1)
	D := make([][]int64, N+1)
	for i := 1; i <= N; i++ {
		F[i] = make([]int64, M+1)
		D[i] = make([]int64, M+1)
	}

	outs := make([][]int, N+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if outs[u] == nil {
			outs[u] = make([]int, 0, 3)
		}
		if outs[v] == nil {
			outs[v] = make([]int, 0, 3)
		}
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	var dfs func(p, u int)
	dfs = func(p, u int) {
		leaf := true
		for _, v := range outs[u] {
			if v == p {
				continue
			}

			leaf = false
			dfs(u, v)
		}

		if leaf {
			for i := 1; i <= M; i++ {
				F[u][i] = 1
			}
		} else {
			for j := 1; j <= M; j++ {
				x := int64(1)
				for _, v := range outs[u] {
					if v == p {
						continue
					}
					x *= D[v][j]
					x %= MOD
				}
				F[u][j] = x
			}
		}

		for i := 1; i <= M; i++ {
			for j := i; j <= M; j += i {
				D[u][i] += F[u][j]
				if D[u][i] >= MOD {
					D[u][i] -= MOD
				}
			}
		}
	}

	dfs(0, 1)

	return int(D[1][1])
}
