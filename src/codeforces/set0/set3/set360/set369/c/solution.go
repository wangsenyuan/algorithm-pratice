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
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	fmt.Println(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int, E [][]int) []int {
	// 就是每条路径上最后一条坏路的候选人要选中
	g := NewGraph(n, 2*n)

	for i, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	var res []int
	// bad, 到u为止，是否有坏的路，返回在子树u中，是否有选中的
	var dfs func(p int, u int, bad bool) bool
	dfs = func(p int, u int, bad bool) bool {
		ret := false
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			j := g.val[i]
			tmp := dfs(u, v, E[j][2] == 2)
			ret = ret || tmp
		}

		if !ret && bad {
			res = append(res, u+1)
			return true
		}

		return ret
	}

	dfs(0, 0, false)

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	g.val = make([]int, e+10)
	// start from 2, reverse of edge i is i ^ 1
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
