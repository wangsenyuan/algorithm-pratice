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
		n, q := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			s, _ := reader.ReadBytes('\n')
			var k int
			pos := readInt(s, 0, &k)
			Q[i] = make([]int, k+1)
			for j := 0; j <= k; j++ {
				pos = readInt(s, pos+1, &Q[i][j])
			}
		}
		res := solve(n, E, Q)
		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
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

func solve(n int, E [][]int, Q [][]int) []int64 {
	cn := make([]int, n)
	for i := 0; i < len(Q); i++ {
		for j := 1; j < len(Q[i]); j++ {
			u := Q[i][j] - 1
			cn[u]++
		}
	}
	qn := make([][]int, n)
	for i := 0; i < n; i++ {
		qn[i] = make([]int, 0, cn[i])
	}

	for i := 0; i < len(Q); i++ {
		for j := 1; j < len(Q[i]); j++ {
			u := Q[i][j] - 1
			qn[u] = append(qn[u], i)
		}
	}
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ind := make([]int, n)
	mp := make([]map[Pair]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make(map[Pair]int)
		ind[i] = i
	}
	ans := make([]int64, len(Q))
	var dfs func(p, u, d int)
	dfs = func(p, u, d int) {
		big := u
		bigSize := 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v, d+1)
			if len(mp[ind[v]]) > bigSize {
				bigSize = len(mp[ind[v]])
				big = v
			}
		}
		ind[u] = ind[big]
		for _, k := range qn[u] {
			dd := Q[k][0]
			ans[k] += int64(mp[ind[u]][Pair{d + dd, k}])
		}
		for _, k := range qn[u] {
			mp[ind[u]][Pair{d, k}]++
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			w := g.to[i]
			if p == w || big == w {
				continue
			}

			for k, v := range mp[ind[w]] {
				dd := Q[k.second][0]
				ans[k.second] += int64(v) * int64(mp[ind[u]][Pair{dd + 2*d - k.first, k.second}])
			}

			for k, v := range mp[ind[w]] {
				mp[ind[u]][k] += v
			}
		}
	}

	dfs(-1, 0, 0)

	return ans
}

type Pair struct {
	first, second int
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
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
