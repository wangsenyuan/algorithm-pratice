package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	t0 := NewDSU(n)
	t1 := NewDSU(n)

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if w == 0 {
			t0.Union(u, v)
		} else {
			t1.Union(u, v)
		}
	}

	var ans int

	for i := 0; i < n; i++ {
		if t0.Find(i) == i {
			sz := t0.cnt[i]
			ans += sz * (sz - 1)
		}
		if t1.Find(i) == i {
			sz := t1.cnt[i]
			ans += sz * (sz - 1)
		}
		// 如果i是那个中点000i111
		x := t0.cnt[t0.Find(i)]
		y := t1.cnt[t1.Find(i)]
		ans += (x - 1) * (y - 1)
	}

	return ans
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

func solve1(n int, edges [][]int) int {
	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	// 0 表示全部是0边的情况, 1表示全部是1边的情况
	// 2 表示先有1边，再有0边的情况
	// 3 表示先有0边，再有1边的情况
	dp := make([][4]int, n+1)

	var ans int

	merge := func(u int, v int, w int) {
		if w == 0 {
			// 起点在u中，必须从0开始
			ans += (1 + dp[u][0]) * (1 + dp[v][1] + dp[v][0] + dp[v][2])
			ans += (1 + dp[v][0]) * (1 + dp[u][1] + dp[u][0] + dp[u][2])
		} else {
			// 如果起点在u中, 可以从0或者1开始，但必须全部是0/1
			// 加1是把u节点算进去
			ans += (dp[u][1] + dp[u][0] + dp[u][3] + 1) * (dp[v][1] + 1)
			// 起点在v中
			ans += (dp[v][1] + dp[v][0] + dp[v][3] + 1) * (dp[u][1] + 1)
		}

		if w == 0 {
			dp[u][0] += dp[v][0] + 1
			dp[u][2] += dp[v][2]
			dp[u][2] += dp[v][1]
		} else {
			dp[u][1] += dp[v][1] + 1
			dp[u][3] += dp[v][3]
			dp[u][3] += dp[v][0]
		}
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				merge(u, v, g.val[i])
			}
		}
	}

	dfs(0, 1)

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
