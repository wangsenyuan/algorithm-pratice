package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	k := readNum(reader)
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res, add := solve(n, edges, queries)

	if len(res) > 0 {
		var buf bytes.Buffer
		buf.WriteString("YES\n")
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", len(x)))
			for i := 0; i < len(x); i++ {
				buf.WriteString(fmt.Sprintf("%d ", x[i]))
			}
			buf.WriteByte('\n')
		}

		fmt.Print(buf.String())
	} else {
		fmt.Println("NO")
		fmt.Println(add)
	}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, edges [][]int, queries [][]int) ([][]int, int) {
	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	fa := make([]int, n)
	depth := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if !vis[v] {
				fa[v] = u
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)
	// deg[u] after all must be even
	deg := make([]int, n)

	process := func(a int, b int) []int {
		// path from a
		var ap []int
		// path from b
		var bp []int

		for a != b {
			if depth[a] >= depth[b] {
				ap = append(ap, a)
				deg[a]++
				deg[fa[a]]++
				a = fa[a]
			} else {
				bp = append(bp, b)
				deg[b]++
				deg[fa[b]]++
				b = fa[b]
			}
		}
		ap = append(ap, a)
		reverse(bp)
		ap = append(ap, bp...)
		return ap
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		ans[i] = process(u, v)
		for j := 0; j < len(ans[i]); j++ {
			ans[i][j]++
		}
	}

	// a -> b -> c -> d
	// 好像不可能出现只有一个奇数节点的情况
	// 因为每次操作至少两个2点，一开始所有的点都是偶数，所以，一次操作后，至少会修改两个点的奇偶性
	// 那么在出现奇数deg的节点时，至少时两个节点
	// 而且可以进一步证明，始终是偶数个奇数节点
	// 所以可以在两个一组间，连接，同时可以保证中间的节点仍然是偶数的，且只改变了端点的奇偶性
	var cnt int
	for i := 0; i < n; i++ {
		cnt += deg[i] & 1
	}

	if cnt > 0 {
		return nil, cnt / 2
	}
	return ans, 0
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
