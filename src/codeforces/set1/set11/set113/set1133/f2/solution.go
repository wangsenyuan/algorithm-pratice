package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, D := readThreeNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges, D)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func solve(n int, edges [][]int, D int) [][]int {
	set := NewDSU(n + 1)

	var keep []int

	g := NewGraph(n+1, 2*n)

	marked := make([]bool, len(edges))
	for i, e := range edges {
		u, v := e[0], e[1]
		if u == 1 || v == 1 {
			keep = append(keep, i)
		} else if set.Union(u, v) {
			marked[i] = true
			g.AddEdge(u, v, i)
			g.AddEdge(v, u, i)
		}
	}
	var deg int
	var keep2 []int
	for _, i := range keep {
		e := edges[i]
		u, v := e[0], e[1]
		if set.Find(u) != set.Find(v) {
			g.AddEdge(u, v, i)
			g.AddEdge(v, u, i)
			marked[i] = true
			deg++
			set.Union(u, v)
		} else {
			keep2 = append(keep2, i)
		}
	}

	// 现在都连在一起了，且在keep2中的边，把它和1连接后，必须删除一条边
	if deg < D {
		// 现在是棵树
		use := make([]int, n+1)
		for i := 0; i <= n; i++ {
			use[i] = -1
		}

		var dfs func(p int, u int)
		dfs = func(p int, u int) {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					use[v] = g.val[i]
					dfs(u, v)
				}
			}
		}

		dfs(0, 1)

		for _, i := range keep2 {
			e := edges[i]
			u := 1 ^ e[0] ^ e[1]
			marked[use[u]] = false
			marked[i] = true
			deg++
			if deg == D {
				break
			}
		}
	}

	if deg != D {
		return nil
	}
	var res [][]int
	for i := range len(edges) {
		if marked[i] {
			res = append(res, edges[i])
		}
	}
	return res
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
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
