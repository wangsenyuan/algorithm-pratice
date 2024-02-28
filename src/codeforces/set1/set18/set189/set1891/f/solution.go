package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		ops := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				var u int
				readInt(s, 2, &u)
				ops[i] = []int{1, u}
			} else {
				var u, x int
				pos := readInt(s, 2, &u)
				readInt(s, pos+1, &x)
				ops[i] = []int{2, u, x}
			}
		}
		res := solve(ops)
		x := fmt.Sprintf("%v", res)
		x = x[1 : len(x)-1]
		buf.WriteString(x)
		buf.WriteByte('\n')
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

func solve(ops [][]int) []int {
	n := 1
	for _, op := range ops {
		if op[0] == 1 {
			n++
		}
	}

	g := NewGraph(n+1, n)
	sz := 1

	id := make([]int, len(ops))

	for i, op := range ops {
		if op[0] == 1 {
			u := op[1]
			sz++
			id[i] = sz
			g.AddEdge(u, sz)
		}
	}

	in := make([]int, n+1)
	out := make([]int, n+1)
	var time int
	var dfs func(u int)
	dfs = func(u int) {
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
		}
		out[u] = time
		time++
	}

	dfs(1)

	ans := make([]int, n+1)

	tr := make(fenwick, time+10)

	for i := len(ops) - 1; i >= 0; i-- {
		cur := ops[i]
		if cur[0] == 1 {
			u := id[i]
			ans[u] = tr.pre(in[u])
		} else {
			u, x := cur[1], cur[2]
			tr.update(in[u], x)
			tr.update(out[u], -x)
			if u == 1 {
				// 节点1比较特殊
				ans[1] += x
			}
		}
	}

	return ans[1:]
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
