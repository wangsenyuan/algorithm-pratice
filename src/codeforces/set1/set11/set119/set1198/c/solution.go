package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		_, _, kind, res := process(reader)
		buf.WriteString(kind)
		buf.WriteByte('\n')
		if kind != imp {
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (n int, edges [][]int, kind string, res []int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	kind, res = solve(n, edges)
	return
}

const mtc = "Matching"
const ind = "IndSet"
const imp = "Impossible"

func solve(n int, edges [][]int) (kind string, res []int) {
	g := NewGraph(3*n, 2*len(edges))
	connect := func(u int, v int, i int) {
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		connect(u, v, i)
	}

	var ans []int

	marked := make([]bool, 3*n)

	used := make([]bool, len(edges))

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		if !marked[u] && !marked[v] {
			ans = append(ans, i+1)
			marked[u] = true
			marked[v] = true
			used[i] = true
		}
	}

	if len(ans) >= n {
		return mtc, ans[:n]
	}

	var nodes []int

	for u := 0; u < 3*n; u++ {
		ok := true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			j := g.val[i]
			if used[j] {
				ok = false
				break
			}
		}
		if ok {
			nodes = append(nodes, u+1)
		}
	}

	return ind, nodes[:n]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
