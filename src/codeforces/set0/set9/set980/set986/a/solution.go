package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	line := readNNums(reader, 4)
	a := readNNums(reader, line[0])
	roads := make([][]int, line[1])
	for i := range line[1] {
		roads[i] = readNNums(reader, 2)
	}
	return solve(line[3], line[2], a, roads)
}

const inf = 1 << 30

func solve(s int, k int, a []int, roads [][]int) []int {
	n := len(a)

	g := NewGraph(n, 2*len(roads))
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)

	bfs := func(x int) []int {
		dist := make([]int, n)

		var head, tail int
		for i := 0; i < n; i++ {
			dist[i] = inf
			if a[i]-1 == x {
				que[head] = i
				head++
				dist[i] = 0
			}
		}

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] == inf {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
		return dist
	}

	dists := make([][]int, k)
	for i := range k {
		dists[i] = bfs(i)
	}

	arr := make([]int, k)

	ans := make([]int, n)

	for u := 0; u < n; u++ {
		for i := 0; i < k; i++ {
			arr[i] = dists[i][u]
		}
		sort.Ints(arr)
		for i := 0; i < s; i++ {
			ans[u] += arr[i]
		}
	}

	return ans
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
