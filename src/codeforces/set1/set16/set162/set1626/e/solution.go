package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := 1

	for tc > 0 {
		tc--
		n := readNum(reader)
		C := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		res := solve(n, C, E)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, C []int, E [][]int) []int {
	// 如果一个节点已经是黑色的，那么它必然是可以到达的
	// 黑色节点的临近节点也没有问题
	// 如果三个以上的黑色节点，是不是一定可以达到呢？
	// 假设某个节点u的临近节点v可以到达黑色节点
	// 且使得v到达黑色节点的第一个黑色节点选择是x，则如果在v的另外一边存在一个y(不同于x）黑色节点,
	// 则u也是可达的
	// 所以，这里首先要知道在v的另外一边有哪些黑色节点可以选择
	// 是不是只要知道在v的另外一边有多少个黑色的节点就可以了？
	// 如果v本身就是黑色的，good， 或者v的另外一边有2个黑色的的节点 good，
	// others bad
	// 现在就是怎么知道v的另外一边有多少个黑色节点呢
	// dfs？
	var blk_cnt int

	for _, x := range C {
		blk_cnt += x
	}

	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	blk := make([]int, n)
	FA := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				FA[v] = u
				dfs(u, v)
				sz[u] += sz[v]
				blk[u] += blk[v]
			}
		}
		sz[u]++
		blk[u] += C[u]
	}

	dfs(0, 0)

	que := make([]int, n)
	var front, end int
	push := func(u int) {
		que[end] = u
		end++
	}
	pop := func() int {
		res := que[front]
		front++
		return res
	}
	vis := make([]int, n)
	for i := 0; i < n; i++ {
		if C[i] == 1 {
			vis[i] = 1
			push(i)
		}
	}

	for front < end {
		u := pop()

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 0 {
				// not vis yet
				if C[u] == 1 {
					vis[v] = 1
					push(v)
				} else {
					// check condition 2
					if FA[v] == u {
						// u is parent
						cnt := blk_cnt - blk[v]
						if cnt >= 2 {
							// 只要能到达u，且使用另外一个黑色节点即可
							vis[v] = 1
							push(v)
						}
					} else {
						// v is parent
						if blk[u] >= 2 {
							vis[v] = 1
							push(v)
						}
					}
				}
			}
		}
	}

	return vis
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
