package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(num int) int {
		fmt.Printf("%d\n", num)
		return readNum(reader)
	}

	n := readNum(reader)
	E := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	solve(n, E, ask)

	fmt.Println("0")

	readNum(reader)
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

func solve(n int, E [][]int, ask func(int) int) {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}
	prev := -1
	curr := -1
	query := func(u int) int {
		prev = curr
		curr = ask(1 + u)
		return curr
	}
	stack := make([]int, 0, n)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
		stack = append(stack, u)
	}

	dfs(-1, 0)

	query(0)

	sameColor := make([]int, n)

	for i := 0; i < n; i++ {
		u := stack[i]
		query(u)
		var same, diff int
		for j := g.nodes[u]; j > 0; j = g.next[j] {
			v := g.to[j]
			if sameColor[v] == 0 {
				continue
			}
			sameColor[v] *= -1
			if sameColor[v] > 0 {
				same++
			} else {
				diff++
			}
		}
		w := (curr - prev + same - diff + 1) / 2
		if w == 0 {
			sameColor[u] = 1
		} else {
			sameColor[u] = -1
		}
	}

	color := make([]int, n)

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if sameColor[v] > 0 {
					color[v] = color[u]
				} else {
					color[v] = 1 - color[u]
				}
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)

	cnt := make([]int, 2)

	for i := 0; i < n; i++ {
		cnt[color[i]]++
	}

	var flip int
	if cnt[1] < cnt[0] {
		flip = 1
	}

	for i := 0; i < n; i++ {
		if color[i] == flip {
			query(i)
		}
	}
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
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
