package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, a, b := readThreeNums(reader)
	p := readNNums(reader, n)
	ok, res := solve(a, b, p)
	if !ok {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func solve(a int, b int, p []int) (bool, []int) {
	pos := make(map[int]int)

	n := len(p)
	N := n * 2

	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}

	g := NewGraph(N, N*4)
	r := NewGraph(N, N*4)

	connect := func(u int, v int) {
		g.AddEdge(u, v)
		g.AddEdge(v^1, u^1)
		r.AddEdge(v, u)
		r.AddEdge(u^1, v^1)
	}

	for _, x := range p {
		i := pos[x]
		y := a - x
		if j, ok := pos[y]; ok {
			connect(2*i, 2*j)
		} else {
			// a - x 不存在，那么i必须在b中
			connect(2*i, 2*i+1)
		}
		z := b - x
		if j, ok := pos[z]; ok {
			connect(2*i+1, 2*j+1)
		} else {
			connect(2*i+1, 2*i)
		}
	}

	vis := make([]bool, N)
	var order []int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < N; i++ {
		if !vis[i] {
			dfs(i)
		}
	}
	comp := make([]int, N)

	var dfs2 func(u int, cid int)
	dfs2 = func(u int, cid int) {
		comp[u] = cid
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if comp[v] == 0 {
				dfs2(v, cid)
			}
		}
	}
	var cid int
	for i := N - 1; i >= 0; i-- {
		if comp[order[i]] == 0 {
			cid++
			dfs2(order[i], cid)
		}
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if comp[2*i] == comp[2*i+1] {
			return false, nil
		}
		if comp[2*i] < comp[2*i+1] {
			ans[i] = 1
		}
	}

	return true, ans
}

func solve1(a int, b int, p []int) (bool, []int) {
	pos := make(map[int]int)

	n := len(p)
	N := n * 2

	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}

	g := NewGraph(N, N*4)
	r := NewGraph(N, N*4)

	connect := func(u int, v int) {
		g.AddEdge(u, v)
		g.AddEdge(v^1, u^1)
		r.AddEdge(v, u)
		r.AddEdge(u^1, v^1)
	}

	for i, x := range p {
		y := a - x
		if j, ok := pos[y]; ok {
			connect(2*i, 2*j)
		} else {
			// a - x 不存在，那么i必须在b中
			connect(2*i, 2*i+1)
		}
		z := b - x
		if j, ok := pos[z]; ok {
			connect(2*i+1, 2*j+1)
		} else {
			connect(2*i+1, 2*i)
		}
	}

	vis := make([]bool, N)
	var order []int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < N; i++ {
		if !vis[i] {
			dfs(i)
		}
	}
	comp := make([]int, N)

	var dfs2 func(u int, cid int)
	dfs2 = func(u int, cid int) {
		comp[u] = cid
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if comp[v] == 0 {
				dfs2(v, cid)
			}
		}
	}
	var cid int
	for i := N - 1; i >= 0; i-- {
		if comp[order[i]] == 0 {
			cid++
			dfs2(order[i], cid)
		}
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if comp[2*i] == comp[2*i+1] {
			return false, nil
		}
		if comp[2*i] < comp[2*i+1] {
			ans[i] = 1
		}
	}

	return true, ans
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
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
