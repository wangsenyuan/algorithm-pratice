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
		ok, res, _ := process(reader)
		if !ok {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
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

func process(reader *bufio.Reader) (ok bool, res []int, mat [][]int) {
	n := readNum(reader)
	mat = make([][]int, 2)
	for i := 0; i < 2; i++ {
		mat[i] = readNNums(reader, n)
	}
	ok, res = solve(mat)
	return
}

func solve(mat [][]int) (bool, []int) {
	n := len(mat[0])

	last := make([]int, n+1)

	g := NewGraph(n, 4*n)

	add := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	cnt := make([]int, n+1)
	for j := 0; j < n; j++ {
		x := mat[0][j]
		y := mat[1][j]
		cnt[x]++
		cnt[y]++
		if x == y {
			continue
		}
		if last[x] != 0 {
			i := last[x]
			if i < 0 {
				// 它在第二行，这两列必须同时交换，或者同时不交换
				add(j, -i-1, 0)
			} else {
				// i > 0, 这两列只能有一列进行交换
				add(j, i-1, 1)
			}
		}
		if last[y] != 0 {
			i := last[y]
			if i < 0 {
				add(j, -i-1, 1)
			} else {
				add(j, i-1, 0)
			}
		}
		last[x] = j + 1
		last[y] = -(j + 1)
	}
	for i := 1; i <= n; i++ {
		if cnt[i] != 2 {
			return false, nil
		}
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}
	comp := make([][]int, 2)
	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		comp[c] = append(comp[c], u+1)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, c^g.val[i]) {
				return false
			}
		}
		return true
	}

	var ans []int

	for i := 0; i < n; i++ {
		if color[i] < 0 {
			clear(comp)
			if !dfs(i, 0) {
				return false, nil
			}
			if len(comp[0]) < len(comp[1]) {
				ans = append(ans, comp[0]...)
			} else {
				ans = append(ans, comp[1]...)
			}
		}
	}

	return true, ans
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
