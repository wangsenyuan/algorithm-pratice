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

	n := readNum(reader)
	nodes := make([][]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = readNNums(reader, 2)
	}

	res := solve(nodes)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(nodes [][]int) []int {
	n := len(nodes)

	var root int

	cnt := make([]int, n)

	g := NewGraph(n, n)

	for i, node := range nodes {
		p, c := node[0], node[1]
		if p == 0 {
			root = i
		} else {
			g.AddEdge(p-1, i)
		}
		cnt[i] = c
	}

	type pair struct {
		first  int
		second int
	}

	sz := make([]int, n)

	dp := make([][]*pair, n)

	merge := func(u, v int) []*pair {
		if sz[u] < sz[v] {
			u, v = v, u
		}
		// sz[u] >= sz[v]
		for _, it := range dp[v] {
			dp[u] = append(dp[u], &pair{it.first + sz[u], it.second})
		}
		return dp[u]
	}

	var dfs func(u int) bool
	dfs = func(u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v) {
				return false
			}
			// 这样可以保证每个节点的值不一样
			dp[u] = merge(u, v)
			sz[u] += sz[v]
		}

		sz[u]++

		if cnt[u] >= sz[u] {
			return false
		}

		if sz[u] == 1 {
			// u赋值为1
			dp[u] = append(dp[u], &pair{1, u})
			return true
		}
		slices.SortFunc(dp[u], func(a, b *pair) int {
			return a.first - b.first
		})
		var val int
		// dp[u]所有的结果的值都必须要不一样
		if cnt[u] < sz[u]-1 {
			it := dp[u][cnt[u]]
			val = it.first
			for i := cnt[u]; i < len(dp[u]); i++ {
				dp[u][i].first++
			}
		} else {
			// 最后一个
			val = dp[u][cnt[u]-1].first + 1
		}
		dp[u] = append(dp[u], &pair{val, u})

		return true
	}

	if !dfs(root) {
		return nil
	}

	ans := make([]int, n)

	for _, it := range dp[root] {
		ans[it.second] = it.first
	}

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
