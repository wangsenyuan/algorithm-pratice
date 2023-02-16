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
		E[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	P := make([][]int, m)
	for i := 0; i < m; i++ {
		P[i] = readNNums(reader, 2)
	}
	res := solve(n, E, P)
	fmt.Println(res)
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

const H = 20

func solve(n int, E [][]int, paths [][]int) int64 {
	g := NewGraph(n+1, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	P := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		P[i] = make([]int, H)
	}
	D := make([]int, n+1)
	var orient func(p, u int)
	orient = func(p, u int) {
		D[u] = D[p] + 1
		P[u][0] = p

		for i := 1; i < H && P[u][i-1] > 0; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				orient(u, v)
			}
		}
	}

	orient(0, 1)

	var ans int64

	var find func(u int)

	S1 := make([]int, n+1)
	S2 := make([]int, n+1)
	Q := make([][]Pair, n+1)

	find = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != P[u][0] {
				find(v)
				S1[u] += S1[v]
				S2[u] += S2[v]
			}
		}
		mem := make(map[Pair]int)
		var sum = S2[u]
		for _, p := range Q[u] {
			sum--
			ans += int64(sum)

			x := getSub(D, P, p.first, u)
			y := getSub(D, P, p.second, u)

			if x != -1 {
				S1[x]--
				ans -= int64(S1[x])
			}
			if y != -1 {
				S1[y]--
				ans -= int64(S1[y])
			}
			if x != -1 && y != -1 {
				ans += int64(mem[Pair{x, y}])
				mem[Pair{x, y}]++
				mem[Pair{y, x}]++
			}
		}
	}

	for i := 0; i < len(paths); i++ {
		u, v := paths[i][0], paths[i][1]
		l := lca(D, P, u, v)
		S1[u]++
		S1[v]++
		S1[l] -= 2
		S2[u]++
		S2[v]++
		S2[l]--
		S2[P[l][0]]--
		if len(Q[l]) == 0 {
			Q[l] = make([]Pair, 0, 2)
		}
		Q[l] = append(Q[l], Pair{u, v})
	}

	find(1)

	return ans
}

type Pair struct {
	first, second int
}

func lca(D []int, P [][]int, u, v int) int {
	if D[u] > D[v] {
		u, v = v, u
	}
	for i := H - 1; i >= 0; i-- {
		if D[v]-(1<<i) >= D[u] {
			v = P[v][i]
		}
	}
	if v == u {
		return u
	}

	for i := H - 1; i >= 0; i-- {
		if P[u][i] != P[v][i] {
			u = P[u][i]
			v = P[v][i]
		}
	}
	return P[u][0]
}

func getSub(D []int, P [][]int, u int, p int) int {
	if u == p {
		return -1
	}
	d := D[u] - D[p] - 1
	for i := H - 1; i >= 0; i-- {
		if (d>>i)&1 == 1 {
			u = P[u][i]
		}
	}
	return u
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
