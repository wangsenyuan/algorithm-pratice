package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, m)

		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)

		if len(res) == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			fmt.Println(res)
		}

	}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, E [][]int) string {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	color := make([]int, n)
	sum := make([][]int, 2)
	for i := 0; i < 2; i++ {
		sum[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	h := NewGraph(n, n)
	var dfs1 func(u int)
	var time int
	var cnt int
	var flip int
	dfs1 = func(u int) {
		in[u] = time
		time++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if color[v] < 0 {
				color[v] = 1 ^ color[u]
				h.AddEdge(u, v)
				dfs1(v)
			} else if in[v] < in[u] {
				dif := color[v] ^ color[u]
				if dif == 0 {
					cnt++
					flip = 1 ^ color[u]
				}
				sum[dif][u]++
				sum[dif][v]--
			}
		}

		out[u] = time
	}

	color[0] = 0
	dfs1(0)

	if cnt <= 1 {
		res := make([]byte, n)
		for i := 0; i < n; i++ {
			res[i] = byte('0' + (flip ^ color[i]))
		}
		return string(res)
	}

	var dfs2 func(u int)

	sv := -1

	dfs2 = func(u int) {
		for i := h.nodes[u]; i > 0; i = h.next[i] {
			v := h.to[i]
			dfs2(v)
			if sum[0][v] == cnt && sum[1][v] == 1 {
				sv = v
				flip = color[u] ^ 1
			}
			for j := 0; j < 2; j++ {
				sum[j][u] += sum[j][v]
			}
		}
	}

	dfs2(0)

	if sv < 0 {
		return ""
	}

	isAnc := func(u int, v int) int {
		if in[u] <= in[v] && out[v] <= out[u] {
			return 1
		}
		return 0
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte('0' + (isAnc(sv, i) ^ flip ^ color[i]))
	}
	return string(res)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) Reset() {
	g.cur = 0
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i] = 0
	}
}
