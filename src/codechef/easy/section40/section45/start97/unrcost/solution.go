package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	q := readNum(reader)

	Q := readNInt64s(reader, q)

	res := solve(n, E, Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func solve(n int, E [][]int, Q []int64) []int {
	g := NewGraph(n, len(E))
	r := NewGraph(n, len(E))
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		r.AddEdge(v, u)
	}

	ord := make([]int, 0, n)
	vis := make([]bool, n)
	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		ord = append(ord, u)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}
	for i := 0; i < n; i++ {
		vis[i] = false
	}
	var dfs2 func(u int) int
	dfs2 = func(u int) int {
		vis[u] = true
		sz := 1
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if !vis[v] {
				sz += dfs2(v)
			}
		}
		return sz
	}
	var scc []int
	for i := n - 1; i >= 0; i-- {
		u := ord[i]
		if !vis[u] {
			scc = append(scc, dfs2(u))
		}
	}

	sort.Ints(scc)
	reverse(scc)

	id := make([]int, len(Q))
	for i := 0; i < len(Q); i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return Q[id[i]] < Q[id[j]]
	})

	ans := make([]int, len(Q))
	var sum = int64(scc[0])
	var cnt int64

	for i, j := 0, 1; i < len(id); i++ {
		k := Q[id[i]]
		for j < len(scc) && cnt+sum*int64(scc[j]) < k {
			cnt += sum * int64(scc[j])
			sum += int64(scc[j])
			j++
		}
		// cnt < k && cnt + sum * int64(scc[j]) > k
		tmp := (k - cnt) / sum
		ans[id[i]] = min(n, int(tmp+sum))
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
