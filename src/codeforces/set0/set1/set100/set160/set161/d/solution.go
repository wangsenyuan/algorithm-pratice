package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, k, E)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, k int, E [][]int) int64 {
	// k <= 500, 可以用dp计算，先用dsu方法计算
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	height := make([]int, n)
	big := make([]int, n)

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		big[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				height[v] = height[u] + 1
				dfs1(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
			}
		}
		sz[u]++
	}

	dfs1(-1, 0)

	var dfs2 func(p int, u int)
	freq := make([]map[int]int, n)
	nodes := make([]map[int]bool, n)
	var ans int64
	dfs2 = func(p int, u int) {

		if height[u] >= k {
			// (u, kth-anc)
			ans++
		}
		freq[u] = make(map[int]int)
		nodes[u] = make(map[int]bool)
		if big[u] > 0 {
			// have children
			//nodes[u] = nodes[big[u]]
			//dfs2(u, big[u], true)
			//nodes[u] = nodes[big[u]]
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				dfs2(u, v)
			}

			freq[u] = freq[big[u]]
			nodes[u] = nodes[big[u]]

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v || big[u] == v {
					continue
				}
				for w := range nodes[v] {
					// h[w] - h[u] + h[x] - h[u]  = k
					hx := k + 2*height[u] - height[w]
					if hx > height[u] {
						ans += int64(freq[u][hx])
					}
				}
				for w := range nodes[v] {
					freq[u][height[w]]++
					nodes[u][w] = true
				}
			}
		}
		freq[u][height[u]]++
		nodes[u][u] = true
	}

	dfs2(-1, 0)

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
