package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q := readTwoNums(reader)
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			var x, v int
			pos := readInt(s, 2, &x)
			readInt(s, pos+1, &v)
			ops[i] = []int{1, x, v}
		} else {
			var x int
			readInt(s, 2, &x)
			ops[i] = []int{2, x}
		}
	}
	res := solve(a, edges, ops)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func solve(a []int, edges [][]int, ops [][]int) []int {
	n := len(a)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	var time int
	parity := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				parity[v] = parity[u] ^ 1
				dfs(u, v)
			}
		}
		out[u] = time
		time++
	}

	dfs(0, 0)

	trs := make([]fenwick, 2)

	for i := 0; i < 2; i++ {
		trs[i] = make(fenwick, time+10)
	}

	var ans []int

	for _, cur := range ops {
		if cur[0] == 1 {
			u := cur[1] - 1
			x := cur[2]
			trs[parity[u]].update(in[u], x)
			trs[parity[u]].update(out[u], -x)
		} else {
			u := cur[1] - 1
			p := parity[u]
			tmp := trs[p].pre(in[u]) - trs[1^p].pre(in[u]) + a[u]
			ans = append(ans, tmp)
		}
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

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
