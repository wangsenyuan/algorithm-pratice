package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	res := solve(n, edges)
	var buf bytes.Buffer
	for i := range n {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i].first, res[i].second))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) []pair {
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vs := make([][]int, n)

	ans := make([]pair, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		var sum int
		bst := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sum += 2 * len(vs[v])
				if bst < 0 || len(vs[v]) > len(vs[bst]) {
					bst = v
				}
			}
		}

		if bst < 0 {
			vs[u] = append(vs[u], u)
			ans[u] = pair{1, 2}
			return
		}

		vs[u] = vs[bst]
		last := ans[bst].second
		sum -= 2 * len(vs[bst])
		sum++
		ans[bst].second += sum

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || v == bst {
				continue
			}

			add := last - 1
			for _, w := range vs[v] {
				ans[w].first += add
				ans[w].second += add
				vs[u] = append(vs[u], w)
			}
			last = ans[v].second
			sum -= 2 * len(vs[v])
			ans[v].second += sum
			clear(vs[v])
		}

		vs[u] = append(vs[u], u)
		ans[u] = pair{last, ans[bst].second + 1}
	}

	dfs(0, 0)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
