package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		p := readNNums(reader, n-1)
		res := solve(n, p)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(n int, p []int) int {
	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		g.AddEdge(p[i]-1, i+1)
	}

	next := make([]int, n)
	var dfs func(u int)
	sz := make([]int, n)
	// dp[i] 表示只考虑在i子树中，处理后，最后能剩下多少个节点未被处理掉的
	dfs = func(u int) {
		next[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)

			if next[u] < 0 || sz[v] > sz[next[u]] {
				next[u] = v
			}
			sz[u] += sz[v]
		}
		sz[u]++
	}

	dfs(0)

	var res int
	var dfs2 func(u int, k int)

	dfs2 = func(u int, k int) {
		if next[u] < 0 {
			return
		}
		v := next[u]
		// already k matched
		if sz[v]-k <= sz[u]-1-sz[v] {
			// 全部可以被match到
			res += (sz[u] - 1 - k) / 2
			return
		}
		// (sz[v] - k) * 2 > sz[u] - 1
		// 这k个全部用来匹配最大的子树
		add := sz[u] - 1 - sz[v]
		res += add

		dfs2(v, max(0, k+add-1))
	}

	dfs2(0, 0)

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
