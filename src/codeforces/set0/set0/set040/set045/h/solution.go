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
		roads[i] = readNNums(reader, 2)
	}

	cnt, res := solve(n, roads)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	if cnt > 0 {
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, edges [][]int) (int, [][]int) {
	if n == 2 {
		// 2个节点，且不能重复边，所以没有答案
		return -1, nil
	}

	marked := make([][]bool, n)
	for i := 0; i < n; i++ {
		marked[i] = make([]bool, n)
	}
	// 不能有bridge， 所有的bridge都需要一条额外的边
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if marked[u][v] {
			continue
		}
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		marked[u][v] = true
		marked[v][u] = true
	}

	low := make([]int, n)
	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = -1
	}

	// 把图变成一颗树
	var comp []int
	belong := make([]int, n)
	stack := make([]int, n)

	var top int

	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = timer
		low[u] = timer
		timer++

		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if in[v] < 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], in[v])
			}
		}

		if in[u] == low[u] {
			comp = append(comp, stack[top-1])
			for top > 0 {
				x := stack[top-1]
				belong[x] = len(comp) - 1
				top--
				if x == u {
					break
				}
			}
		}
	}

	dfs(-1, 0)

	// comp组成一颗新的树
	g2 := make([]map[int]bool, len(comp))

	for i := 0; i < len(g2); i++ {
		g2[i] = make(map[int]bool)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u = belong[u]
		v = belong[v]
		if u != v {
			g2[u][v] = true
			g2[v][u] = true
		}
	}

	var arr []int

	for i := 0; i < len(g2); i++ {
		if len(g2[i]) == 1 {
			arr = append(arr, comp[i])
		}
	}

	var res [][]int

	for i := 0; i < (len(arr)+1)/2; i++ {
		u := arr[i]
		v := arr[min(len(arr)-1, i+(len(arr)+1)/2)]
		res = append(res, []int{u + 1, v + 1})
	}

	return len(res), res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, 3+e)
	to := make([]int, 3+e)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
