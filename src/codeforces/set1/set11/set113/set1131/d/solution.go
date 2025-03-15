package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	adj := make([]string, n)
	for i := range n {
		adj[i] = readString(reader)[:m]
	}

	x, row, col := solve(adj)

	if x < 0 {
		fmt.Println("No")
		return
	}
	var buf bytes.Buffer

	buf.WriteString("Yes\n")
	for _, x := range row {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	for _, x := range col {
		buf.WriteString(fmt.Sprintf("%d ", x))
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
func solve(adj []string) (int, []int, []int) {
	n := len(adj)
	m := len(adj[0])

	// 需要把相等的那些节点，变成一个节点
	comp := NewDSU(n + m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if adj[i][j] == '=' {
				comp.Union(i, n+j)
			}
		}
	}

	g := NewGraph(n+m, 2*n*m)

	deg := make([]int, n+m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := comp.Find(i)
			y := comp.Find(n + j)
			if (adj[i][j] == '=') != (x == y) {
				return -1, nil, nil
			}

			if adj[i][j] == '=' {
				continue
			}

			if adj[i][j] == '>' {
				x, y = y, x
			}
			g.AddEdge(x, y)
			deg[y]++
		}
	}
	// 必须从头开始
	dp := make([]int, n+m)
	que := make([]int, n+m)
	var head, tail int
	for i := 0; i < n+m; i++ {
		dp[i] = -1
		j := comp.Find(i)
		if i == j && deg[i] == 0 {
			que[head] = i
			head++
			dp[i] = 1
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dp[v] = max(dp[v], dp[u]+1)
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	// 如果有cycle => false
	for i := 0; i < n+m; i++ {
		j := comp.Find(i)
		if deg[j] > 0 {
			return -1, nil, nil
		}
	}

	row := make([]int, n)
	col := make([]int, m)

	for i := 0; i < n; i++ {
		row[i] = dp[comp.Find(i)]
	}

	for i := 0; i < m; i++ {
		col[i] = dp[comp.Find(i+n)]
	}

	return slices.Max(dp), row, col
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
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
