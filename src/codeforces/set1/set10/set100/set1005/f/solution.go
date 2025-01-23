package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (n int, edges [][]int, res []string) {
	n, m, k := readThreeNums(reader)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(n, k, edges)
	return
}

func solve(n int, k int, edges [][]int) (s []string) {
	m := len(edges)

	g := NewGraph(n, 2*m)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	dist := make([]int, n)
	for i := range n {
		dist[i] = -1
	}
	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	dist[0] = 0

	buf := make([]byte, m)
	for i := range m {
		buf[i] = '0'
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	use := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		if dist[u] > dist[v] {
			u, v = v, u
		}
		if dist[u]+1 == dist[v] {
			use[v] = append(use[v], i)
		}
	}
	pos := make([]int, n)

	for range k {
		for i := range m {
			buf[i] = '0'
		}
		for i := 1; i < n; i++ {
			buf[use[i][pos[i]]] = '1'
		}
		s = append(s, string(buf))
		ok := false
		for i := 1; i < n; i++ {
			if pos[i]+1 < len(use[i]) {
				ok = true
				pos[i]++
				break
			} else {
				pos[i] = 0
			}
		}
		if !ok {
			break
		}
	}

	return
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.val = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
