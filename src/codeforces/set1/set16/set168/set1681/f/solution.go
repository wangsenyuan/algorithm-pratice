package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, E [][]int) int64 {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		w--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	sz := make([]int, n)
	pw := make([]int, n)
	ord := make([]int, 0, n)
	in := make([]int, n)
	out := make([]int, n)
	var dfs1 func(p int, u int)
	var time int
	dfs1 = func(p int, u int) {
		in[u] = time
		time++
		ord = append(ord, u)
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				pw[v] = g.val[i]
				dfs1(u, v)
				sz[u] += sz[v]
			}
		}
		out[u] = time
	}

	dfs1(0, 0)

	vs := make([][]int, n)

	for _, u := range ord {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			w := g.val[i]
			if vs[w] == nil {
				vs[w] = make([]int, 0, 1)
				vs[w] = append(vs[w], 0)
			}
			vs[w] = append(vs[w], u)
		}
	}

	isAnc := func(u int, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	h := NewGraph(n, n)

	var dfs2 func(v int, x int) int64

	dfs2 = func(v int, x int) int64 {
		var res int64
		dp[v][0] = sz[v]

		for i := h.nodes[v]; i > 0; i = h.next[i] {
			u := h.to[i]
			res += dfs2(u, x)
			dp[v][0] -= sz[u]
		}
		dp[v][1] = 0
		for i := h.nodes[v]; i > 0; i = h.next[i] {
			u := h.to[i]
			if pw[u] == x {
				res += int64(dp[u][0]) * int64(dp[v][0])
				dp[v][1] += dp[u][0]
			} else {
				res += int64(dp[u][0]) * int64(dp[v][1])
				res += int64(dp[u][1]) * int64(dp[v][0])
				dp[v][0] += dp[u][0]
				dp[v][1] += dp[u][1]
			}
		}
		return res
	}

	stack := make([]int, n)
	var ans int64
	for x := 0; x < n; x++ {
		if len(vs[x]) == 0 {
			continue
		}
		vs[x] = unique(vs[x])
		var p int
		for _, v := range vs[x] {
			for p > 0 && !isAnc(stack[p-1], v) {
				p--
			}
			if p > 0 {
				h.AddEdge(stack[p-1], v, 0)
			}
			stack[p] = v
			p++
		}
		ans += dfs2(0, x)
		h.cur = 0
		for _, v := range vs[x] {
			h.nodes[v] = 0
		}
	}

	return ans
}

func unique(arr []int) []int {
	occ := make(map[int]bool)
	n := len(arr)
	var p int
	for i := 0; i < n; i++ {
		if occ[arr[i]] {
			continue
		}
		occ[arr[i]] = true
		arr[p] = arr[i]
		p++
	}
	return arr[:p]
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

func (g *Graph) Reset(n int) {
	g.cur = 0
	for i := 0; i < n; i++ {
		g.nodes[i] = 0
	}
}
