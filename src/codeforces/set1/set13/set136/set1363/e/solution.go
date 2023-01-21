package main

import (
	"bufio"
	"fmt"
	"math"
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

	n := readNum(reader)

	X := make([][]int, n)
	for i := 0; i < n; i++ {
		X[i] = readNNums(reader, 3)
	}
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	fmt.Println(solve(n, X, E))
}

func solve(n int, X [][]int, E [][]int) int64 {
	adj := getAdj(n, E)
	has := make([][]int, n)
	need := make([][]int, n)
	for i := 0; i < n; i++ {
		has[i] = make([]int, 2)
		need[i] = make([]int, 2)
	}

	var dfs1 func(p, u int)
	dfs1 = func(p, u int) {

		if X[u][1] < X[u][2] {
			has[u][0]++
			need[u][1]++
		} else if X[u][1] > X[u][2] {
			has[u][1]++
			need[u][0]++
		}

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs1(u, v)
			for j := 0; j <= 1; j++ {
				has[u][j] += has[v][j]
				need[u][j] += need[v][j]
			}
		}
	}

	dfs1(-1, 0)

	if has[0][0] != need[0][0] || has[0][1] != need[0][1] {
		return -1
	}

	var res int64
	var dfs2 func(p, u, x int) int

	dfs2 = func(p, u, x int) int {
		var handled int
		for _, v := range adj[u] {
			if p == v {
				continue
			}
			handled += dfs2(u, v, min(x, X[u][0]))
		}
		if X[u][0] < x {
			// 当前节点是最小的cost，把剩余没有处理的节点处理掉
			// has[0], has[1], need[0], need[1]
			// min(has[0], need[0]) + min(has[1], need[1])
			y := min(has[u][0], need[u][0]) + min(has[u][1], need[u][1]) - handled
			res += int64(y) * int64(X[u][0])
			handled += y
		}
		return handled
	}

	dfs2(-1, 0, math.MaxInt32)

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func getAdj(n int, E [][]int) [][]int {
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
