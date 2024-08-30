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
	cnt := readNNums(reader, 3)

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges, cnt)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d", res[i]))
	}
	fmt.Println(buf.String())

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

func solve(n int, edges [][]int, expect []int) []int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var dfs func(u int, c int, cnt []int) bool

	dfs = func(u int, c int, cnt []int) bool {
		if color[u] != -1 {
			return color[u] == c
		}
		cnt[c]++
		color[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, c^1, cnt) {
				return false
			}
		}
		return true
	}
	// dp[0][0] = true
	var comp [][]int
	var prev [][]int
	dp := []bool{true}

	var sz int
	for i := 0; i < n; i++ {
		if color[i] < 0 {
			cnt := make([]int, 2)
			if !dfs(i, 0, cnt) {
				return nil
			}
			sz += cnt[0] + cnt[1]
			fp := make([]bool, sz+1)
			from := make([]int, sz+1)

			for j := 0; j < len(dp); j++ {
				if dp[j] {
					fp[j+cnt[0]] = true
					from[j+cnt[0]] = j
					fp[j+cnt[1]] = true
					from[j+cnt[1]] = j
				}
			}
			comp = append(comp, []int{cnt[0], cnt[1], i})
			prev = append(prev, from)
			dp = fp
		}
	}
	n2 := expect[1]

	if n2 >= len(dp) || !dp[n2] {
		return nil
	}

	label := make([]int, n)

	vis := make([]bool, n)

	var dfs2 func(u int, x int)

	dfs2 = func(u int, x int) {
		if color[u] == x {
			label[u] = 2
		}
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs2(v, x)
			}
		}
	}

	for it := len(comp); it > 0; it-- {
		cnt := comp[it-1]
		if n2 >= cnt[0] && n2 < len(prev[it-1]) && n2-prev[it-1][n2] == cnt[0] {
			dfs2(cnt[2], 0)
			n2 -= cnt[0]
		} else {
			dfs2(cnt[2], 1)
			n2 -= cnt[1]
		}
	}

	n1 := expect[0]
	for i := 0; i < n; i++ {
		if label[i] == 2 {
			continue
		}
		if n1 > 0 {
			label[i] = 1
			n1--
		} else {
			label[i] = 3
		}
	}

	return label
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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
