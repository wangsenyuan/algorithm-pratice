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
		reqs := make([][]int, m)
		for i := 0; i < m; i++ {
			reqs[i] = readNNums(reader, 2)
		}
		res := solve(n, reqs)
		if len(res) == 0 {
			buf.WriteString("INFINITE\n")
		} else {
			buf.WriteString(fmt.Sprintf("FINITE\n%d\n", len(res)))
			for _, e := range res {
				buf.WriteString(fmt.Sprintf("%d ", e))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, requirements [][]int) []int {
	if n == 1 {
		return []int{1}
	}
	if len(requirements) == 0 {
		// 无限长
		return nil
	}

	g := NewGraph(n+1, len(requirements))

	for _, cur := range requirements {
		a, b := cur[0], cur[1]
		// from 1 to a
		g.AddEdge(b, a)
	}
	// 如果和1直接相连的节点，它们只能有两个，且在1的两边
	// 假设其中的一个节点是u
	// 那么和u相连的节点, 假设是v，它理论上可以有3个
	// v u 1 v u v
	// 和v相连的节点w
	// w v u 1 w v w u v w
	// 如何优雅的生成这个列表呢？
	// 这里还有一个限制，就是那些len(comp) > 1 的部分，只能出现一次
	// 比如 （3， 2）（2， 3）（3， 1）
	// 2, 3, 1, 2, 3, 2
	// 完蛋了，这个例子表面前面的分析都是错的
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = -1
	}

	que := make([]int, n)
	var front, tail int
	dist[1] = 0
	que[front] = 1
	front++

	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[front] = v
				front++
			}
		}
	}
	var m int
	for i := 2; i <= n; i++ {
		if dist[i] < 0 {
			// no connected to 1
			return nil
		}
		m = max(dist[i], m)
	}

	comps := make([][]int, m+1)
	for i := 1; i <= n; i++ {
		d := dist[i]
		if comps[d] == nil {
			comps[d] = make([]int, 0, 1)
		}
		comps[d] = append(comps[d], i)
	}

	var res []int

	for d := m; d >= 0; d-- {
		for x := d; x <= m; x++ {
			res = append(res, comps[x]...)
		}
	}

	return res
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
