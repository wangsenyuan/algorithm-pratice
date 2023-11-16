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

	n, m, maxDist := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	res := solve(n, edges, maxDist)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	s := fmt.Sprintf("%v", res)
	buf.WriteString(s[1 : len(s)-1])
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

const inf = 1 << 60

func solve(n int, edges [][]int, maxDist int) []int {
	m := len(edges)
	g := NewGraph(n, m)
	rg := NewGraph(n, m)

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		rg.AddEdge(v, u, w)
	}

	stack := make([]int, n)
	// a dag
	var top int
	var dfs func(u int)
	vis := make([]bool, n)
	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		stack[top] = u
		top++
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}
	// stack[top] should be 0
	for top > 0 && stack[top-1] != 0 {
		top--
	}

	// dp[u][i] = 在节点u，通过i个节点到达u时，的最短路径
	dp := make([][]int, n)
	dp[0] = make([]int, 1)
	dp[0][0] = 0
	for it := top - 2; it >= 0; it-- {
		u := stack[it]
		var d int
		for i := rg.nodes[u]; i > 0; i = rg.next[i] {
			v := rg.to[i]
			d = max(d, len(dp[v]))
		}
		d++
		dp[u] = make([]int, d)
		for i := 0; i < d; i++ {
			dp[u][i] = inf
		}
		for i := rg.nodes[u]; i > 0; i = rg.next[i] {
			v := rg.to[i]
			for j := 0; j < len(dp[v]); j++ {
				dp[u][j+1] = min(dp[u][j+1], dp[v][j]+rg.val[i])
			}
		}
	}

	k := len(dp[n-1])

	for k > 0 && dp[n-1][k-1] > maxDist {
		k--
	}
	// k > 0
	ans := make([]int, k)
	k--
	ans[k] = n
	cur := n - 1
	for cur > 0 {
		for i := rg.nodes[cur]; i > 0; i = rg.next[i] {
			v := rg.to[i]
			w := rg.val[i]
			if len(dp[v]) >= k && dp[v][k-1]+w == dp[cur][k] {
				ans[k-1] = v + 1
				cur = v
				break
			}
		}

		k--
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
