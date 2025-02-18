package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for range tc {
		readString(reader)
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, m, edges)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(num ...int) int {
	r := 1
	for _, x := range num {
		r = r * x % mod
	}
	return r
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, k int, edges [][]int) int {
	if k == 2 {
		return mul(n, n-1, pow(2, mod-2))
	}
	deg := make([]int, n)

	g := NewGraph(n, n*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	// k > 2
	dist := make([]int, n)
	que := make([]int, n)
	from := make([]int, n)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, n)
	}

	dp := make([]int, k+1)

	bfs := func(x int) int {
		if deg[x] < k {
			// no contribution
			return 0
		}
		for i := range n {
			dist[i] = -1
			from[i] = -1
			clear(cnt[i])
		}

		dist[x] = 0
		var head, tail int

		for i := g.nodes[x]; i > 0; i = g.next[i] {
			v := g.to[i]
			dist[v] = 1
			from[v] = v
			cnt[v][1]++
			que[head] = v
			head++
		}

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					from[v] = from[u]
					cnt[from[v]][dist[v]]++
					que[head] = v
					head++
				}
			}
		}

		var res int

		for d := n - 1; d > 0; d-- {
			clear(dp)
			// 找到距离为d的k个节点
			dp[0] = 1
			var w int
			for i := g.nodes[x]; i > 0; i = g.next[i] {
				v := g.to[i]
				c := cnt[v][d]
				w++
				if c > 0 {
					// 这里怎么感觉出现了k这个系数呢
					for x := min(w, k); x > 0; x-- {
						dp[x] = add(dp[x], mul(dp[x-1], c))
					}
				}
			}

			res = add(res, dp[k])
		}

		return res
	}

	var res int
	for u := 0; u < n; u++ {
		res = add(res, bfs(u))
	}

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
