package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, k, E)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func solve(n int, rem int, E [][]int) bool {
	g := NewGraph(2*(n+1), 2*len(E))

	for _, e := range E {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}
	pr := make([]int, 2*(n+1))
	pl := make([]int, n+1)
	var argument func(u int) bool

	vis := make([]int, 2*(n+1))
	var time int

	argument = func(u int) bool {
		if vis[u] == time {
			return false
		}
		vis[u] = time
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if pr[v] == 0 || argument(pr[v]) {
				pr[v] = u
				pl[u] = v
				return true
			}
		}
		return false
	}

	for u := 1; u <= n; u++ {
		time++
		if !argument(u) {
			return false
		}
	}

	var dfs func()
	var cycle func(int)

	frozen := make([]bool, n+1)
	in := make([]bool, n+1)
	erased := make([]bool, 2*len(E)+3)
	fa := make([]int, n+1)
	var root int
	ok := false

	erase := func(x int, ok bool) {
		for i := g.node[x]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == pl[x] {
				erased[i] = ok
				return
			}
		}
	}

	dfs = func() {
		if rem == 0 {
			return
		}

		p := make([]int, n+1)
		copy(p, pl)
		time++
		ok = false
		for i := 1; i <= n && !ok; i++ {
			if !frozen[i] && vis[i] != time {
				cycle(i)
			}
		}
		if !ok {
			return
		}

		x := root
		for {
			p[x] = fa[x]
			x = pr[fa[x]]
			if x == root {
				break
			}
		}
		frozen[x] = true
		dfs()
		frozen[x] = false
		if rem == 0 {
			return
		}
		rem--
		if rem == 0 {
			return
		}
		erase(x, true)

		for i := 1; i <= n; i++ {
			p[i], pl[i] = pl[i], p[i]
			pr[pl[i]] = i
		}

		dfs()

		for i := 1; i <= n; i++ {
			p[i], pl[i] = pl[i], p[i]
			pr[pl[i]] = i
		}

		erase(x, false)
	}

	cycle = func(x int) {
		if ok {
			return
		}
		vis[x] = time
		in[x] = true
		y := pl[x]
		for i := g.node[y]; i > 0; i = g.next[i] {
			if ok {
				break
			}
			z := g.to[i]
			// i ^ 1 is the edge from z to y
			if !frozen[z] && !erased[i^1] && z != x {
				fa[z] = y
				if vis[z] != time {
					cycle(z)
				} else if in[z] {
					// a cycle found
					ok = true
					root = z
				}
			}
		}
		in[x] = false
	}

	dfs()

	return rem == 0
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
