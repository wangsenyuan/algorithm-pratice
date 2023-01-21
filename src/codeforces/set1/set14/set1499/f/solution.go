package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, k, E)
	fmt.Println(res)
}

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

const MOD = 998244353

func modAdd(a, b int64) int64 {
	a += b
	a %= MOD
	return a
}

func modMul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}

func solve(n, k int, E [][]int) int {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, n+1)
	}

	var dfs func(p, u int) int

	aux := make([]int64, n+1)

	dfs = func(p, u int) int {
		dp[u][0] = 1
		h := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			vh := dfs(u, v)
			uh := max(h, vh+1)

			for a := 0; a < uh; a++ {
				aux[a] = 0
			}

			for a := 0; a < h; a++ {
				for b := 0; b < vh; b++ {
					if a+b+1 <= k {
						aux[max(a, b+1)] = modAdd(aux[max(a, b+1)], modMul(dp[u][a], dp[v][b]))
					}
					if a <= k && b <= k {
						aux[a] = modAdd(aux[a], modMul(dp[u][a], dp[v][b]))
					}
				}
			}
			h = uh
			for a := 0; a < h; a++ {
				dp[u][a] = aux[a]
			}
		}
		return h
	}
	dfs(-1, 0)

	var res int64
	for i := 0; i <= k; i++ {
		res = modAdd(res, dp[0][i])
	}

	return int(res)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
