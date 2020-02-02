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
		n := line[0]
		k := line[2]
		A := readNNums(scanner, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		res := solve(n, k, A, edges)
		fmt.Println(res)
	}
}

const INF = 1 << 26

func solve(n int, k int, A []int, edges [][]int) int {
	K := 1 << uint(k)

	dp := make([][]map[int]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]map[int]int, 2)
		//dp[i][0] = not select i, dp[i][1] = select i
		dp[i][0] = make(map[int]int)
		dp[i][1] = make(map[int]int)

	}

	conns := make([][]int, n)
	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	arr := make([]int, 2*n+3)

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		d := 1
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs1(u, v)
			arr[2*u+d] = v
			d++
		}
	}

	dfs1(-1, 0)

	for i := 0; i < n; i++ {
		A[i]--
	}

	var dfs2 func(u int)

	dfs2 = func(u int) {
		left := arr[2*u+1]
		right := arr[2*u+2]

		if left == 0 && right == 0 {
			// it is leaf
			dp[u][0][0] = 0
			dp[u][1][1<<uint(A[u])] = 1
			return
		}

		if left > 0 {
			dfs2(left)
		}

		if right > 0 {
			dfs2(right)
		}

		// it is internal node
		// try not select u

		if left > 0 && right > 0 {
			// have two outs
			// try select u, then we have option not to select left & right
			for k1, c1 := range dp[left][0] {
				for k2, c2 := range dp[right][0] {
					k := k1 | k2 | (1 << uint(A[u]))
					if c, found := dp[u][1][k]; !found || c1+c2+1 < c {
						dp[u][1][k] = c1 + c2 + 1
					}
				}

				for k2, c2 := range dp[right][1] {
					k := k1 | k2 | (1 << uint(A[u]))
					if c, found := dp[u][1][k]; !found || c1+c2+1 < c {
						dp[u][1][k] = c1 + c2 + 1
					}
				}
			}

			for k1, c1 := range dp[left][1] {
				for k2, c2 := range dp[right][0] {
					k := k1 | k2 | (1 << uint(A[u]))
					if c, found := dp[u][1][k]; !found || c1+c2+1 < c {
						dp[u][1][k] = c1 + c2 + 1
					}
				}

				for k2, c2 := range dp[right][1] {
					k := k1 | k2

					if c, found := dp[u][0][k]; !found || c1+c2 < c {
						dp[u][0][k] = c1 + c2
					}

					k |= (1 << uint(A[u]))
					if c, found := dp[u][1][k]; !found || c1+c2+1 < c {
						dp[u][1][k] = c1 + c2 + 1
					}
				}
			}
			return
		}
		if left > 0 {
			for k1, c1 := range dp[left][0] {
				k := k1 | (1 << uint(A[u]))

				if c, found := dp[u][1][k]; !found || c1+1 < c {
					dp[u][1][k] = c1 + 1
				}
			}

			for k1, c1 := range dp[left][1] {
				k := k1

				if c, found := dp[u][0][k]; !found || c1+1 < c {
					dp[u][0][k] = c1
				}

				k |= (1 << uint(A[u]))
				if c, found := dp[u][1][k]; !found || c1+1 < c {
					dp[u][1][k] = c1 + 1
				}
			}
		}

		// no sigle right child
	}

	dfs2(0)

	res1 := getResult(dp[0][0], K-1, INF)
	res2 := getResult(dp[0][1], K-1, INF)
	res := min(res1, res2)
	if res >= INF {
		return -1
	}
	return res
}

func getResult(mem map[int]int, k int, defalutValue int) int {
	if v, found := mem[k]; found {
		return v
	}
	return defalutValue
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
