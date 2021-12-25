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

	tc :=readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		A := readNInt64Nums(scanner, n)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, A, m, edges))
	}
}

const INF = 1 << 50
func solve(n int, A []int64, M int, edges [][]int) int {
	outs := make([][]int, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0] - 1, edge[1] - 1
		if A[v] > A[u] {
			outs[u] = append(outs[u], v)
		} else if A[u] > A[v] {
			outs[v] = append(outs[v], u)
		}
	}

	var dfs func(u int)
	V := make([]int, n)

	dfs = func(u int) {
		V[u] = 1

		for _, v := range outs[u] {
			dfs(v)

			if len(outs[v]) == 0 {
				V[u] = max(V[u], 2)
				continue
			}
			for _, w := range outs[v] {
				if A[w] - A[v] > A[v] - A[u] {
					V[u] = max(V[u], V[w] + 2)
				}
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		if V[i] == 0 {
			dfs(i)
			res = max(res, V[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
