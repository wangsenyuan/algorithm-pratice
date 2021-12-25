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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, X := readTwoNums(scanner)
		A := readNNums(scanner, N)
		edges := make([][]int, N - 1)
		for i := 0; i < N - 1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(N, X, A, edges))
	}
}

func solve(N int, X int, A []int, edges [][]int) int {

	outs := make([][]int, N)
	for i := 0; i < N; i++ {
		outs[i] = make([]int,0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0] - 1, edge[1] - 1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	var dfs func(p, u int) int

	dfs = func(p, u int) int {
		var res int

		for _, v := range outs[u] {
			if v == p {
				continue
			}
			res += dfs(u, v)

		}
		return max(res + A[u], -X)
	}

	return dfs(-1, 0)
}


func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}