package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 2)
	}
	k := readNum(reader)
	path := readNNums(reader, k)

	res := solve(n, roads, path)

	fmt.Println(res[0], res[1])
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

const inf = 1 << 30

func solve(n int, roads [][]int, path []int) []int {
	g := NewGraph(n, len(roads))
	r := NewGraph(n, len(roads))
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(v, u)
		r.AddEdge(u, v)
	}

	for i := 0; i < len(path); i++ {
		path[i]--
	}

	t := path[len(path)-1]

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	dist[t] = 0

	que := make([]int, n)
	var head, tail int
	que[head] = t
	head++

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	dp := make([]int, n)
	fp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
		fp[i] = -inf
	}

	res := []int{0, 0}

	for i := 0; i+1 < len(path); i++ {
		u := path[i]
		var cnt int
		found := false

		for j := r.nodes[u]; j > 0; j = r.next[j] {
			v := r.to[j]
			if dist[u] == dist[v]+1 {
				cnt++
				if v == path[i+1] {
					found = true
				}
			}
		}

		if !found {
			// 要rebuild一次
			res[0]++
			res[1]++
			continue
		}
		if cnt > 1 {
			//除了v还有个w
			res[1]++
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
