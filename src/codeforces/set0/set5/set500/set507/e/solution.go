package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}

	res := solve(n, edges)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
	}

	fmt.Print(buf.String())
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

const inf = 1 << 60

func solve(n int, edges [][]int) [][]int {
	g := NewGraph(n, len(edges)*2)

	var tot_working int

	for i, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
		tot_working += w
	}

	que := make([]int, n)

	bfs := func(x int) []int {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}

		var head, tail int
		que[head] = x
		head++
		dist[x] = 0

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}

		return dist
	}

	d0 := bfs(0)
	d1 := bfs(n - 1)

	m := d0[n-1]

	type pair struct {
		first  int
		second int
	}

	var arr []pair
	ok := make([]bool, n)
	for i := 0; i < n; i++ {
		if d0[i]+d1[i] == m {
			arr = append(arr, pair{d0[i], i})
			ok[i] = true
		}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})
	// dp[i]表示那些最短路径上的节点i，能够使用的最多的working的edge数目
	dp := make([]int, n)
	from := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = -inf
		from[i] = -1
	}

	dp[0] = 0
	for i := 1; i < len(arr); i++ {
		u := arr[i].second
		for j := g.nodes[u]; j > 0; j = g.next[j] {
			v := g.to[j]
			eid := g.val[j]
			z := edges[eid][2]
			if ok[v] && d0[v]+1 == d0[u] {
				if dp[v]+z > dp[u] {
					dp[u] = dp[v] + z
					from[u] = eid
				}
			}
		}
	}

	var ans [][]int
	// 先找出那些需要
	use := make([]bool, len(edges))

	for i := n - 1; ; {
		eid := from[i]
		use[eid] = true
		edge := edges[eid]
		u, v, z := edge[0]-1, edge[1]-1, edge[2]
		j := u ^ v ^ i

		if z == 0 {
			// need to repair
			ans = append(ans, []int{i + 1, j + 1, 1})
		}
		i = j
		if i == 0 {
			break
		}
	}

	for i, e := range edges {
		if use[i] {
			continue
		}
		if e[2] == 1 {
			// need to blow up
			ans = append(ans, []int{e[0], e[1], 0})
		}
	}

	return ans
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
