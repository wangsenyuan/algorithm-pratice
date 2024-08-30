package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}

	res := solve(n, edges)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%.10f ", res[i]))
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

func abs(num int) int {
	return max(num, -num)
}

const inf = 1 << 50

type pair struct {
	first  int
	second int
}

const eps = 1e-10

func solve(n int, edges [][]int) []float64 {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	vis := make([]bool, n)
	val := make([]pair, n)

	stack := make([]int, n)

	var top int

	var x *pair

	var dfs func(p int, u int) bool

	dfs = func(p int, u int) bool {
		vis[u] = true
		stack[top] = u
		top++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			tmp := pair{val[u].first * -1, g.val[i] - val[u].second}
			if !vis[v] {
				val[v] = tmp
				if !dfs(u, v) {
					return false
				}
				continue
			}
			if val[v].first == tmp.first {
				if val[v].second != tmp.second {
					return false
				}
				continue
			}
			// val[v].first * tmp.first = -1
			tx := pair{tmp.first - val[v].first, val[v].second - tmp.second}
			if x == nil {
				x = &tx
			} else if x.first*tx.second != x.second*tx.first {
				return false
			}
		}
		return true
	}

	f := func(v float64) float64 {
		var res float64
		for i := 0; i < top; i++ {
			u := stack[i]
			res += math.Abs(v*float64(val[u].first) + float64(val[u].second))
		}
		return res
	}

	ans := make([]float64, n)

	for i := 0; i < n; i++ {
		if !vis[i] {
			val[i] = pair{1, 0}
			top = 0
			x = nil
			if !dfs(-1, i) {
				return nil
			}
			var xv float64
			if x != nil {
				xv = float64(x.second) / float64(x.first)
			} else {
				var l, r float64 = -100000000, 100000000
				for math.Abs(r-l) > eps {
					m1 := l + (r-l)/3
					m2 := r - (r-l)/3
					if f(m1) > f(m2) {
						l = m1
					} else {
						r = m2
					}
				}
				xv = (l + r) / 2
			}
			for i := 0; i < top; i++ {
				u := stack[i]
				ans[u] = xv*float64(val[u].first) + float64(val[u].second)
			}
		}
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
