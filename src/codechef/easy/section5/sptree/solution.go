package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n, K, a := readThreeNums(reader)
		F := readNNums(reader, K)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, a, E, F)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i].first))
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i].second))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(n int, r int, E [][]int, F []int) []Pair {
	//以r为根，计算d(r,u) - d(x,u)最大值
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	// S[i] 表示以i为根的子树中是否有特殊节点
	S := make([]int, n)
	X := make([]bool, n)
	for _, x := range F {
		S[x-1] = x
		X[x-1] = true
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p {
				continue
			}
			D[v] = D[u] + 1
			dfs(u, v)
			if S[v] > 0 {
				// for u, it doesn't matter which sub-node is special
				S[u] = S[v]
			}
		}
	}

	dfs(-1, r-1)

	ans := make([]Pair, n)

	var dfs2 func(p, u int, x Pair)

	dfs2 = func(p, u int, x Pair) {
		if S[u] > 0 {
			// u的子树中有特殊节点
			ans[u] = Pair{D[u], S[u]}
		} else {
			// 在r 到 u 的path中，存在一个特殊节点
			f, s := x.first, x.second
			if f >= 0 {
				ans[u] = Pair{D[f] - (D[u] - D[f]), s}
			} else {
				ans[u] = Pair{-D[u], s}
			}
		}

		var first = -1
		var second = -1

		if X[u] {
			first = u
		}

		for i := g.nodes[u]; i > 0 && second < 0; i = g.next[i] {
			v := g.to[i]
			if p != v && S[v] > 0 {
				// some special nodes in S[v]
				if first < 0 {
					first = v
				} else if second < 0 {
					second = v
				}
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if first >= 0 && first != v {
					dfs2(u, v, Pair{u, S[first]})
				} else if second >= 0 && second != v {
					dfs2(u, v, Pair{u, S[second]})
				} else {
					dfs2(u, v, x)
				}
			}
		}
	}

	dfs2(-1, r-1, Pair{-1, F[0]})

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
