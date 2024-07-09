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
		n := readNum(reader)
		reqs := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var cnt int
			pos := readInt(s, 0, &cnt) + 1
			reqs[i] = make([]int, cnt)
			for j := 0; j < cnt; j++ {
				pos = readInt(s, pos, &reqs[i][j]) + 1
			}
		}
		res := solve(n, reqs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, requests [][]int) int {
	deg := make([]int, n)
	var sum int
	for u, cur := range requests {
		deg[u] = len(cur)
		sum += len(cur)
	}

	g := NewGraph(n, sum)

	for u, cur := range requests {
		for _, v := range cur {
			g.AddEdge(v-1, u)
		}
	}

	dist := make([]int, n)
	que := make([]int, n*2)
	front, end := n, n
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[front] = i
			front++
		}
	}

	for end < front {
		u := que[end]
		end++

		for j := g.nodes[u]; j > 0; j = g.next[j] {
			v := g.to[j]
			deg[v]--
			if u < v {
				dist[v] = max(dist[v], dist[u])
			} else {
				// u > v
				dist[v] = max(dist[v], dist[u]+1)
			}

			if deg[v] == 0 {
				if dist[v] == dist[u] {
					end--
					que[end] = v
				} else {
					que[front] = v
					front++
				}
			}
		}
	}
	var ans int

	for i := 0; i < n; i++ {
		if deg[i] > 0 {
			return -1
		}
		ans = max(ans, dist[i])
	}

	return ans + 1
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
