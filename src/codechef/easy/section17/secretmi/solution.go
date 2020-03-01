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

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		line := readNNums(scanner, 3)
		N, M, L := line[0], line[1], line[2]
		B := readNNums(scanner, L)
		edges := make([][]int, M)
		for i := 0; i < M; i++ {
			edges[i] = readNNums(scanner, 3)
		}
		fmt.Println(solve(N, M, L, B, edges))
	}
}

func solve(N, M, L int, B []int, edges [][]int) int {
	for i := 0; i < L; i++ {
		B[i]--
	}

	// then we need to find min distance between each pair of cities (i, j)
	fp := make([][]int, N)
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		fp[i] = make([]int, N)
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
		dp[i][i] = 0
	}

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		dp[u][v] = w
		dp[v][u] = w
		fp[u][v] = w
		fp[v][u] = w
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			if dp[i][k] < 0 {
				continue
			}
			for j := 0; j < N; j++ {
				if dp[k][j] < 0 {
					continue
				}
				if dp[i][j] < 0 || dp[i][j] > dp[i][k]+dp[k][j] {
					dp[i][j] = dp[i][k] + dp[k][j]
				}
			}
		}
	}

	res := 1
	prev := 0

	var dist int

	for i := 1; i < L; i++ {
		dist += fp[B[i-1]][B[i]]
		if dist != dp[B[prev]][B[i]] {
			//there have to be a target city between (prev, i)
			if i-prev == 1 || dp[B[i-1]][B[i]] != fp[B[i-1]][B[i]] {
				return -1
			}
			res++
			prev = i - 1
			dist = fp[B[i-1]][B[i]]
		}
	}

	if dist != dp[B[prev]][B[L-1]] {
		return -1
	}
	res++
	return res
}
