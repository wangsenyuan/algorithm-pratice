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
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	res := solve(n, edges)
	if len(res) == 0 {
		fmt.Println("NO\n")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func solve(n int, E [][]int) []int {
	g := NewGraph(n+1, 2*n)

	for _, e := range E {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	fa := make([]int, n+1)
	dist := make([]int, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				fa[v] = u
				dfs(u, v)
			}
		}
	}

	dfs(0, 1)

	first := 1
	for i := 1; i <= n; i++ {
		if dist[i] > dist[first] {
			first = i
		}
	}

	dist[first] = 0
	fa[first] = 0
	dfs(0, first)

	second := 1
	for i := 1; i <= n; i++ {
		if dist[i] > dist[second] {
			second = i
		}
	}
	ok := make([]bool, n+1)
	var dia []int
	x := second
	for x != 0 {
		ok[x] = true
		dia = append(dia, x)
		x = fa[x]
	}

	for i := 1; i <= n; i++ {
		if !ok[i] && !ok[fa[i]] {
			// no solution
			return nil
		}
	}

	var res []int

	for i := 0; i < len(dia); i += 2 {
		res = append(res, dia[i])
		if i+1 < len(dia) {
			u := dia[i+1]
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				if !ok[v] {
					res = append(res, v)
				}
			}
		}
	}

	sub := len(dia) & 1

	for i := len(dia) - 1; i > 0; i -= 2 {
		res = append(res, dia[i-sub])
		if i-1-sub >= 0 {
			u := dia[i-1-sub]
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				if !ok[v] {
					res = append(res, v)
				}
			}
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
