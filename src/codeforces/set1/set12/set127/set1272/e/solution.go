package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(a []int) []int {
	n := len(a)

	x := make([]int, n+1)
	copy(x[1:], a)
	a = x

	g := make([]*Graph, 2)
	for i := 0; i < 2; i++ {
		g[i] = NewGraph(n+1, 2*n)
	}

	arr := make([][]int, 2)

	for i := 1; i <= n; i++ {
		ok := false
		if i-a[i] >= 1 {
			j := i - a[i]
			if a[i]&1 == a[j]&1 {
				g[a[i]&1].AddEdge(j, i)
			} else {
				// 从i可以1次到不同parity的位置
				ok = true
			}
		}
		if i+a[i] <= n {
			j := i + a[i]
			if a[i]&1 == a[j]&1 {
				g[a[i]&1].AddEdge(j, i)
			} else {
				ok = true
			}
		}
		if ok {
			arr[a[i]&1] = append(arr[a[i]&1], i)
		}
	}

	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = -1
	}
	que := make([]int, n)

	bfs := func(arr []int, g *Graph) {
		var head, tail int
		for _, x := range arr {
			dist[x] = 1
			que[head] = x
			head++
		}

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
	}

	bfs(arr[0], g[0])
	bfs(arr[1], g[1])

	return dist[1:]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
