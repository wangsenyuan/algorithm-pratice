package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
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

const MOD = 19101995

func solve(n int, E [][]int) int {
	g := NewGraph(n, 2*len(E))
	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	var lg int
	for 1<<uint(lg) <= n {
		lg++
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, lg+1)
		for j := 0; j <= lg; j++ {
			dp[i][j] = -1
		}
		dp[i][lg] = 0
	}

	var dfs func(p, u, cnt int)

	dfs = func(p, u, cnt int) {
		if dp[u][cnt] >= 0 {
			return
		}

		sz := g.size[u]

		if p >= 0 {
			sz--
		}
		if sz == 0 {
			dp[u][cnt] = 1
			return
		}
		children := make([]int, sz)
		for i, j := g.nodes[u], 0; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			children[j] = v
			j++
			// pick edge(u, v) as light
			dfs(u, v, cnt+1)
		}

		left := make([]int64, sz)
		dp[u][cnt] = 0

		for j := 0; j < sz; j++ {
			left[j] = int64(dp[children[j]][cnt+1])
			if j > 0 {
				left[j] *= left[j-1]
				left[j] %= MOD
			}
		}
		var right int64 = 1
		for j := sz - 1; j >= 0; j-- {
			//pick edge(u, v) as heavy
			dfs(u, children[j], cnt)
			tmp := right * int64(dp[children[j]][cnt]) % MOD
			if j > 0 {
				tmp *= left[j-1]
				tmp %= MOD
			}
			dp[u][cnt] += int(tmp)
			dp[u][cnt] %= MOD
			right = right * int64(dp[children[j]][cnt+1]) % MOD
		}
	}

	dfs(-1, 0, 0)

	return dp[0][0]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	size  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.size = make([]int, n)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.size[u]++
}
