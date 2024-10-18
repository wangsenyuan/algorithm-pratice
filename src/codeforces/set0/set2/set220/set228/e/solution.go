package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	roads := make([][]int, m)

	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 3)
	}

	ok, res := solve(n, roads)

	if !ok {
		fmt.Println("Impossible")
		return
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, roads [][]int) (bool, []int) {
	N := n * 2

	g := NewGraph(N, len(roads)*10)
	connect := func(u int, v int) {
		g.AddEdge(u, v)
		g.AddEdge(v^1, u^1)
	}

	for _, road := range roads {
		u, v, w := road[0]-1, road[1]-1, road[2]
		if w == 1 {
			// 两个必须同时为true
			connect(2*u, 2*v)
			// 或者同时为false
			connect(2*u+1, 2*v+1)
		} else {
			connect(2*u, 2*v+1)
			connect(2*u+1, 2*v)
		}
	}

	low := make([]int, N)
	dsc := make([]int, N)
	vis := make([]int, N)
	comp := make([]int, N)
	var cid int
	stack := make([]int, N)
	var top int
	var time int

	var dfs func(u int)
	dfs = func(u int) {
		dsc[u] = time
		low[u] = time
		time++
		stack[top] = u
		top++
		vis[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 2 {
				continue
			}
			if vis[v] == 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], dsc[v])
			}
		}

		if low[u] == dsc[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				comp[v] = cid
				if u == v {
					break
				}
			}
			cid++
		}
		vis[u]++
	}

	for i := 0; i < N; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}

	var ans []int

	for i := 0; i < n; i++ {
		if comp[2*i] == comp[2*i+1] {
			return false, nil
		}
		if comp[2*i] < comp[2*i+1] {
			ans = append(ans, i+1)
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
