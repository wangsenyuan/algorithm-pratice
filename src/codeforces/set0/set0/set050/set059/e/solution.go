package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _, _ := process(reader)
	fmt.Println(len(res) - 1)
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
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

func process(reader *bufio.Reader) (res []int, n int, roads [][]int, forbiden [][]int) {
	n, m, k := readThreeNums(reader)
	roads = make([][]int, m)
	for i := range m {
		roads[i] = readNNums(reader, 2)
	}
	forbiden = make([][]int, k)
	for i := range k {
		forbiden[i] = readNNums(reader, 3)
	}
	res = solve(n, roads, forbiden)
	return
}

const inf = 1 << 60

type record struct {
	a int
	b int
	c int
}

func solve(n int, roads [][]int, forbiden [][]int) []int {

	bad := make(map[record]bool)
	for _, cur := range forbiden {
		r := record{cur[0] - 1, cur[1] - 1, cur[2] - 1}
		bad[r] = true
	}

	g := NewGraph(n, len(roads)*2)

	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = inf
		}
	}
	que := make([]int, n*n)
	var head, tail int
	// 起始位置
	for i := g.nodes[0]; i > 0; i = g.next[i] {
		v := g.to[i]
		dp[0][v] = 1
		que[head] = v
		head++
	}

	for tail < head {
		a, b := que[tail]/n, que[tail]%n
		tail++
		for i := g.nodes[b]; i > 0; i = g.next[i] {
			c := g.to[i]
			if !bad[record{a, b, c}] && dp[b][c] > dp[a][b]+1 {
				dp[b][c] = dp[a][b] + 1
				que[head] = b*n + c
				head++
			}
		}
	}

	res := inf
	for i := g.nodes[n-1]; i > 0; i = g.next[i] {
		v := g.to[i]
		res = min(res, dp[v][n-1])
	}

	if res == inf {
		return nil
	}
	var path []int
	cur := n - 1

	for cur != 0 {
		path = append(path, cur+1)
		var prev int
		for i := g.nodes[cur]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dp[v][cur] == res {
				prev = v
				break
			}
		}
		res--
		cur = prev
	}
	path = append(path, 1)

	reverse(path)

	return path
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
