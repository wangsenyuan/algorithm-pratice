package main

import (
	"bufio"
	"bytes"
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
	A := readNNums(reader, n)
	m := readNum(reader)
	commands := make([][]int, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == 'F' {
			var u, v int
			pos := readInt(s, 2, &u)
			readInt(s, pos+1, &v)
			commands[i] = []int{u, v}
		} else {
			var u, v, d int
			pos := readInt(s, 2, &u)
			pos = readInt(s, pos+1, &v)
			readInt(s, pos+1, &d)
			commands[i] = []int{u, v, d}
		}
	}

	res := solve(n, E, A, commands)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

func solve(n int, E [][]int, A []int, commands [][]int) []int {
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	big := make([]int, n)

	P := make([][]int, n)
	D := make([]int, n)
	var dfs1 func(p int, u int)
	dfs1 = func(p int, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		big[u] = -1
		sz[u]++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs1(u, v)
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
				sz[u] += sz[v]
			}
		}
	}

	dfs1(0, 0)
	// 节点u所属的chain
	var chains [][]int
	// chain的头部节点
	chainIndex := make([]int, n)
	// 节点u的位置
	chainPos := make([]int, n)
	var num int
	var hld func(u int)

	hld = func(u int) {
		if len(chains) == num {
			chains = append(chains, make([]int, 0, 1))
		}
		chainPos[u] = len(chains[num])
		chains[num] = append(chains[num], A[u])
		chainIndex[u] = num

		if big[u] >= 0 {
			hld(big[u])
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == P[u][0] || v == big[u] {
				continue
			}
			// a new chain
			num++
			hld(v)
		}
	}

	hld(0)
	num++
	trees := make([]*Segment, num)

	for i := 0; i < num; i++ {
		trees[i] = NewSegment(len(chains[i]), chains[i])
	}

	lca := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
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

	find := func(p int, u int) int {
		var ans int

		for {
			if chainIndex[u] == chainIndex[p] {
				ans = gcd(ans, trees[chainIndex[u]].Query(chainPos[p], chainPos[u]))
				break
			}
			ans = gcd(ans, trees[chainIndex[u]].Query(0, chainPos[u]))
			u = P[u][0]
		}
		return ans
	}

	update := func(p int, u int, d int) {
		for {
			if chainIndex[u] == chainIndex[p] {
				trees[chainIndex[u]].Update(chainPos[p], chainPos[u], d)
				break
			}
			trees[chainIndex[u]].Update(0, chainPos[u], d)
			u = P[u][0]
		}
	}

	var res []int

	for _, cmd := range commands {
		if len(cmd) == 2 {
			// query
			u, v := cmd[0], cmd[1]
			l := lca(u, v)

			res = append(res, gcd(find(l, u), find(l, v)))
		} else {
			u, v, d := cmd[0], cmd[1], cmd[2]
			l := lca(u, v)
			update(l, u, d)
			update(l, v, d)
			update(l, l, -d)
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

type Node struct {
	first   int
	last    int
	diffGcd int
	lazy    int
}

func merge(a, b, c *Node) {
	a.first = b.first
	a.last = c.last
	a.diffGcd = gcd(b.diffGcd, gcd(c.first-b.last, c.diffGcd))
}

type Segment struct {
	nodes []*Node
	n     int
}

func NewSegment(n int, arr []int) *Segment {
	nodes := make([]*Node, 4*n)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = new(Node)
	}

	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l == r {
			nodes[i].first = arr[l]
			nodes[i].last = arr[l]
			nodes[i].lazy = 0
			return
		}
		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		merge(nodes[i], nodes[2*i+1], nodes[2*i+2])
	}

	loop(0, 0, n-1)

	return &Segment{nodes, n}
}

func (t *Segment) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			t.nodes[i].first += v
			t.nodes[i].last += v
			t.nodes[i].lazy += v
			return
		}
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		merge(t.nodes[i], t.nodes[2*i+1], t.nodes[2*i+2])
		// t.nodes[i].lazy 没有变化，但是first & last 变化了
		t.nodes[i].first += t.nodes[i].lazy
		t.nodes[i].last += t.nodes[i].lazy
	}

	loop(0, 0, t.n-1)
}

func (t *Segment) Query(L int, R int) int {
	var loop func(i int, l int, r int, sum int) int
	loop = func(i int, l int, r int, sum int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return gcd(t.nodes[i].first+sum, t.nodes[i].diffGcd)
		}
		sum += t.nodes[i].lazy
		mid := (l + r) / 2
		a := loop(2*i+1, l, mid, sum)
		b := loop(2*i+2, mid+1, r, sum)
		return gcd(a, b)
	}
	return loop(0, 0, t.n-1, 0)
}
