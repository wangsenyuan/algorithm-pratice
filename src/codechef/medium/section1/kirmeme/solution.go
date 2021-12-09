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

	for tc > 0 {
		tc--
		n := readNum(reader)
		L, R := readTwoNums(reader)
		A := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, A, L, R)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, E [][]int, A []int, L, R int) int64 {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	del := make([]bool, n)

	size := make([]int, n)
	W := make([]int, n+10)
	var wn int

	var dfsSize func(p, u int)

	dfsSize = func(p, u int) {
		wn++
		W[wn] = u
		size[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !del[v] {
				dfsSize(u, v)
				size[u] += size[v]
			}
		}
	}

	getCentroid := func() int {
		opt := 1 << 30
		var ret int
		for i := 1; i <= wn; i++ {
			k := W[i]
			cur := wn - size[k]
			for j := g.nodes[k]; j > 0; j = g.next[j] {
				v := g.to[j]
				if del[v] || size[v] > size[k] {
					continue
				}
				cur = max(cur, size[v])
			}

			if cur < opt {
				// 最大子树最小的节点
				opt = cur
				ret = k
			}
		}
		return ret
	}

	F := make([][]int64, 2)

	for i := 0; i < 2; i++ {
		F[i] = make([]int64, n+10)
	}

	update := func(i int, p int, v int) {
		p++
		for p <= wn+1 {
			F[i][p] += int64(v)
			p += p & -p
		}
	}

	get := func(i int, p int) int64 {
		p++
		var res int64
		for p > 0 {
			res += F[i][p]
			p -= p & -p
		}
		return res
	}

	getRange := func(i int, l int, r int) int64 {
		res := get(i, min(r, wn))
		res -= get(i, min(wn, l-1))
		return res
	}

	var result int64
	pa := make([]int, n+1)
	var pn int

	var dfs func(p, u int, pks int, tp int)

	dfs = func(p, u int, pks int, tp int) {
		if pks > R {
			return
		}

		if L <= pks {
			result++
		}

		result += getRange(0, max(0, L-pks), R-pks)

		if pks+tp <= R {
			result += getRange(1, max(0, L-pks-tp), R-pks-tp)
		}

		pn++
		pa[pn] = pks

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || del[v] {
				continue
			}
			npks := pks
			if A[u] > A[v] && A[u] > A[p] {
				npks++
			}
			dfs(u, v, npks, tp)
		}

	}

	var loop func(k int)

	loop = func(k int) {
		wn = 0
		dfsSize(-1, k)
		c := getCentroid()
		del[c] = true

		for i := 0; i <= wn+1; i++ {
			F[0][i] = 0
			F[1][i] = 0
		}

		for i := g.nodes[c]; i > 0; i = g.next[i] {
			v := g.to[i]
			if del[v] {
				continue
			}
			var tp int
			if A[v] < A[c] {
				tp = 1
			}
			pn = 0
			dfs(c, v, 0, tp)
			for j := 1; j <= pn; j++ {
				update(tp, pa[j], 1)
			}
		}

		for i := 0; i <= wn+1; i++ {
			F[0][i] = 0
			F[1][i] = 0
		}

		for i := g.nodes[c]; i > 0; i = g.next[i] {
			v := g.to[i]
			if del[v] {
				continue
			}
			loop(v)
		}
	}

	loop(0)

	return result
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
