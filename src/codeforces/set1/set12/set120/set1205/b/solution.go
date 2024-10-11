package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	fmt.Println(res)
}

const H = 60

const inf = 1 << H

func solve(a []int) int {
	// 每个节点顶多60条边，引出去就可以了？
	arr := make([][]int, H)

	var nums []int

	for _, num := range a {
		if num != 0 {
			nums = append(nums, num)
		}
	}

	a = nums

	for id, num := range a {
		for i := 0; i < H; i++ {
			if (num>>i)&1 == 1 {
				arr[i] = append(arr[i], id)
			}
		}
	}

	for i := 0; i < H; i++ {
		if len(arr[i]) >= 3 {
			return 3
		}
	}
	// n不会很大

	best := inf

	n := len(a)

	g := NewGraph(n, 2*H)

	for _, v := range arr {
		// len(v) <= 2
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				g.AddEdge(v[i], v[j])
				g.AddEdge(v[j], v[i])
			}
		}
	}
	que := make([]int, n)
	dist := make([]int, n)
	from := make([]int, n)

	bfs := func(x int) {
		for i := 0; i < n; i++ {
			dist[i] = -1
			from[i] = -1
		}
		dist[x] = 0
		var head, tail int
		que[head] = x
		head++

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					from[v] = u
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				} else if from[u] != v && from[v] != u {
					best = min(best, dist[u]+dist[v]+1)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		bfs(i)
	}

	if best == inf {
		return -1
	}

	return best
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
