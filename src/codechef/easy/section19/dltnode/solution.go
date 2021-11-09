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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		values := readNNums(reader, n)

		res := solve(values, E)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return s[:len(s)-1]
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

func solve(values []int, E [][]int) int64 {
	n := len(values)

	// gcd(child-tree) is ok
	// gcd(self with children-tree) is also ok
	// gcd(parent) ?
	g := NewGraph(n, 2*(n-1))

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	C := make([]int, n)

	in := make([]int, n)
	out := make([]int, n)
	ord := make([]int, n)
	var time int

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		in[u] = time
		ord[time] = u
		time++
		C[u] = values[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				C[u] = gcd(C[u], C[v])
			}
		}
		out[u] = time - 1
	}
	dfs1(0, 0)

	forward := make([]int, time)

	for i := 0; i < time; i++ {
		forward[i] = values[ord[i]]
		if i > 0 {
			forward[i] = gcd(forward[i], forward[i-1])
		}
	}

	backward := make([]int, time)
	for i := time - 1; i >= 0; i-- {
		backward[i] = values[ord[i]]
		if i+1 < time {
			backward[i] = gcd(backward[i], backward[i+1])
		}
	}

	var best int64

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		// if we delete u
		var x int
		if in[u] > 0 {
			x = gcd(x, forward[in[u]-1])
		}
		if out[u]+1 < time {
			x = gcd(x, backward[out[u]+1])
		}

		tmp := int64(x)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
				tmp += int64(C[v])
			}
		}
		if tmp > best {
			best = tmp
		}
	}

	dfs2(0, 0)

	return best
}

func solve1(values []int, E [][]int) int64 {
	n := len(values)

	// gcd(child-tree) is ok
	// gcd(self with children-tree) is also ok
	// gcd(parent) ?
	g := NewGraph(n, 2*(n-1))

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	C := make([]int, n)

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		C[u] = values[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				C[u] = gcd(C[u], C[v])
			}
		}
	}
	dfs1(0, 0)

	var best int64

	var dfs2 func(p, u int, x int)

	arr := make([]int, n)
	dfs2 = func(p, u int, x int) {
		// if we pick u
		tmp := int64(x)
		var it int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp += int64(C[v])
				arr[it] = v
				it++
			}
		}
		if tmp > best {
			best = tmp
		}
		if it == 0 {
			return
		}
		L := make([]int, it)
		R := make([]int, it)
		L[0] = C[arr[0]]

		for i := 1; i < it; i++ {
			L[i] = gcd(L[i-1], C[arr[i]])
		}

		R[it-1] = C[arr[it-1]]

		for i := it - 2; i >= 0; i-- {
			R[i] = gcd(R[i+1], C[arr[i]])
		}

		var ii int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				var y = gcd(x, values[u])
				if ii > 0 {
					y = gcd(y, L[ii-1])
				}
				if ii+1 < it {
					y = gcd(y, R[ii+1])
				}
				dfs2(u, v, y)
				ii++
			}
		}
	}

	dfs2(0, 0, 0)

	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
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
