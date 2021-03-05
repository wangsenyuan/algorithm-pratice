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
		n, m := readTwoNums(reader)
		P := make([][]int, m)
		for i := 0; i < m; i++ {
			P[i] = readNNums(reader, 2)
		}
		res := solve(n, m, P)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
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

func solve(n int, m int, P [][]int) []int {
	g := NewGraph(2*n, m*2)
	deg := make([]int, 2*n)
	id := make(map[[2]int]int)
	var start int
	for i := 0; i < m; i++ {
		r, c := P[i][0], P[i][1]
		r--
		c--
		id[[...]int{r, n + c}] = i
		id[[...]int{n + c, r}] = i
		g.AddEdge(r, n+c, i)
		g.AddEdge(n+c, r, i)
		deg[r] ^= 1
		deg[n+c] ^= 1
	}

	res := make([]int, 0, m+1)
	var dfs func(u int)
	vis := make([]bool, m)
	pos := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		pos[i] = g.nodes[i]
		if deg[i] == 1 {
			start = i
		}
	}
	dfs = func(u int) {
		for pos[u] > 0 {
			v := g.to[pos[u]]
			j := g.val[pos[u]]
			pos[u] = g.next[pos[u]]
			if !vis[j] {
				vis[j] = true
				dfs(v)
			}
		}
		res = append(res, u)
	}
	dfs(start)

	ans := make([]int, 0, m)

	for i := 1; i < len(res); i++ {
		ans = append(ans, id[[...]int{res[i-1], res[i]}]+1)
	}
	return ans
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
