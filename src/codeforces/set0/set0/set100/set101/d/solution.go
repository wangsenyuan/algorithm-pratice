package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	fmt.Printf("%.7f\n", res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, E [][]int) float64 {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	T := make([]int, n)
	sz := make([]int, n)
	dp := make([]int64, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		var nodes [][]int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				dp[u] += dp[v] + int64(g.value[i])*int64(sz[v])
				T[u] += T[v] + 2*g.value[i]
				T[v] += 2 * g.value[i]
				nodes = append(nodes, []int{T[v], sz[v], v})
			}
		}
		sz[u]++
		// sort by T/sz
		sort.Slice(nodes, func(i, j int) bool {
			t1 := nodes[i][0]
			s1 := nodes[i][1]
			t2 := nodes[j][0]
			s2 := nodes[j][1]
			// t1/sz
			return int64(t1)*int64(s2) <= int64(t2)*int64(s1)
		})

		var sum int64

		for i := 0; i < len(nodes); i++ {
			v := nodes[i][2]
			dp[u] += sum * int64(sz[v])
			sum += int64(T[v])
		}
	}

	dfs(0, 0)

	return float64(dp[0]) / float64(n-1)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	value []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.value = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.value[g.cur] = w
}
