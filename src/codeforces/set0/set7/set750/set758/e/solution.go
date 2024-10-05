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
		edges[i] = readNNums(reader, 4)
	}
	res := solve(n, edges)
	if n > 1 && len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", n))
	for _, e := range res {
		for i := 0; i < 4; i++ {
			buf.WriteString(fmt.Sprintf("%d ", e[i]))
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

func solve(n int, edges [][]int) [][]int {

	g := NewGraph(n, n)

	for i, cur := range edges {
		u, v := cur[0], cur[1]
		u--
		v--
		g.AddEdge(u, v, i)
	}

	type record struct {
		minSum   int
		extraDec int
	}

	a := make([]record, n)

	var dfs func(u int) (mn int, mx int)

	dfs = func(u int) (mn int, mx int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			v1, v2 := dfs(v)
			edge := edges[g.val[i]]
			weight, strength := edge[2], edge[3]
			if v1 < 0 || strength < v1 {
				return -1, 0
			}
			// u-v 边的强度可以减少到子树 w 的最小重量和
			mn += max(1, weight-(strength-v1)) + v1
			// 子树 w 的最大重量和不能超过 v-w 边的强度
			mx += weight + min(v2, strength)
			a[v].extraDec += max(0, v2-strength)
		}
		a[u].minSum = mn
		return
	}

	mn, _ := dfs(0)
	if mn < 0 {
		return nil
	}

	res := make([][]int, n-1)
	copy(res, edges)

	var dec int

	var modify func(u int)

	modify = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]

			dec += a[v].extraDec

			modify(v)

			e := res[g.val[i]]
			d := min(e[2]-1, e[3]-a[v].minSum, dec)
			e[2] -= d
			e[3] -= d
			dec -= d
		}
	}

	modify(0)

	return res
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
