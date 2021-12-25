package main

import (
	"bufio"
	"bytes"
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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		N, Q := readTwoNums(reader)
		A := readNNums(reader, N)
		E := make([][]int, N-1)
		for i := 0; i < N-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		V := make([]int, Q)
		for i := 0; i < Q; i++ {
			V[i] = readNum(reader)
		}

		res := solve(N, Q, A, E, V)

		var buf bytes.Buffer

		for i := 0; i < N; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		fmt.Println(buf.String())
	}
}

func solve(N int, Q int, A []int, E [][]int, V []int) []int64 {
	adj := getGraph(N, E)

	X := make([][]int64, N)
	for i := 0; i < N; i++ {
		X[i] = make([]int64, 2)
	}

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		X[u][0] = int64(A[u])

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs1(u, v)
			X[u][1] += X[v][0]
			X[u][0] += X[v][1]
		}
	}

	dfs1(-1, 0)

	Y := make([]bool, N)

	for i := 0; i < Q; i++ {
		Y[V[i]-1] = true
	}

	B := make([]int64, N)

	var dfs2 func(p, u int, flag int, d int)

	dfs2 = func(p, u int, flag int, d int) {
		var cur int = flag
		if d == 0 {
			// even level
			if flag&1 == 1 {
				// even parent visited
				B[u] = 0
			} else if Y[u] {
				B[u] = X[u][0]
				cur |= 1
			} else {
				B[u] = int64(A[u])
			}
		} else {
			if flag&2 == 2 {
				// odd parent visited
				B[u] = 0
			} else if Y[u] {
				B[u] = X[u][0]
				cur |= 2
			} else {
				B[u] = int64(A[u])
			}
		}

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs2(u, v, cur, 1-d)
		}
	}

	dfs2(-1, 0, 0, 0)

	return B
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
	}
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}
