package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(a, edges)
}

func solve(a []int, edges [][]int) []int {
	n := len(a)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = n * (n - 1) / 2
	}
	cnt := make([]int, n+1)
	var all int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1
		old := cnt[a[u]]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				pre := all - cnt[a[u]]
				dfs(u, v)
				sz[u] += sz[v]
				nxt := all - cnt[a[u]]
				ans[a[u]] -= (nxt - pre) * (nxt - pre - 1) / 2
			}
		}
		// u子树中具体的颜色c的分布没有关系， 因为这个他们在上层被计算是，通过减法被抵消掉了
		// all - cnt[c] (all中会计算一次，cnt[c]也被计算了，所以抵消掉了)
		// 但是它的上层，到这里为止，需要计算正确
		cnt[a[u]] = old + sz[u]
		all++
	}

	dfs(-1, 0)

	for c := 1; c <= n; c++ {
		pre := 0
		nxt := n - cnt[c]
		ans[c] -= (nxt - pre) * (nxt - pre - 1) / 2
	}
	for _, c := range a {
		ans[c]++
	}

	return ans[1:]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
