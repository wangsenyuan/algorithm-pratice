package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	T := readString(reader)
	n := readNum(reader)
	S := make([]string, n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		s := readString(reader)
		var j int
		for s[j] != ' ' {
			j++
		}
		S[i] = s[:j]
		readInt([]byte(s[j+1:]), 0, &A[i])
	}
	res := solve(T, S, A)
	fmt.Println(res)
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

const INF = 1 << 50

func solve(T string, S []string, A []int) int64 {
	n := len(S)
	// n <= 100
	// 无法直接证明是否能够获得T，当前行某一个字符的选择，会影响到另外一个字符的选择，因为总数不能超过A[i]
	// 这个是个图的问题
	// T中的字符和包含这个字符的行i，连接，流量是1
	// 行i和sink连接， 流量是A[i]
	// source 和 T中的字符x连接, 流量是freq[x]
	// 最大流
	// 但是cost怎么算呢？
	// 是不是在选取边的时候，优先选择节点小的边即可？
	freq := getFreq(T)
	nodes := 1 + n + 26 + 2

	dist := make([]int64, nodes)

	edges := 26 + 26*n + n
	var src int
	sink := nodes - 1
	g := NewGraph(nodes, edges*2)

	add := func(u, v, c, x int) {
		g.AddEdge(u, v, 0, c, x)
		g.AddEdge(v, u, 0, 0, -x)
	}

	// 从后往前，优先选择小标号的节点
	for i := n - 1; i >= 0; i-- {
		add(src, i+1, A[i], i+1)
	}

	for i := 0; i < n; i++ {
		cnt := getFreq(S[i])
		for j := 0; j < 26; j++ {
			if cnt[j] > 0 && freq[j] > 0 {
				add(i+1, j+(n+1), min(cnt[j], freq[j]), 0)
			}
		}
	}

	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			add(i+(n+1), sink, freq[i], 0)
		}
	}

	en := make([]bool, nodes)
	P := make([]int, nodes)
	X := make([]int, nodes)
	bfs := func() bool {
		for i := 0; i < nodes; i++ {
			dist[i] = INF
			en[i] = false
		}
		var que []int
		que = append(que, src)
		dist[src] = 0
		var front int
		for front < len(que) {
			u := que[front]
			front++
			en[u] = false
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] > dist[u]+int64(g.cost[i]) && g.val[i] < g.lim[i] {
					dist[v] = dist[u] + int64(g.cost[i])
					P[v] = u
					X[v] = i
					if !en[v] {
						en[v] = true
						que = append(que, v)
					}
				}
			}
		}

		return dist[sink] < INF
	}

	var res int64

	var tot int
	for bfs() {
		u := sink
		flow := 1 << 30
		for u != src {
			i := X[u]
			flow = min(flow, g.lim[i]-g.val[i])
			u = P[u]
		}
		tot += flow
		res += int64(flow) * dist[sink]
		u = sink
		for u != src {
			i := X[u]
			g.val[i] += flow
			g.val[i^1] -= flow
			u = P[u]
		}
	}

	if tot < len(T) {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	lim   []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+2)
	to := make([]int, e+2)
	val := make([]int, e+2)
	lim := make([]int, e+2)
	cost := make([]int, e+2)
	// 从1开始，1, 2是互为反方向的边
	return &Graph{nodes, next, to, val, lim, cost, 1}
}

func (g *Graph) AddEdge(u, v, w, c, x int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
	g.lim[g.cur] = c
	g.cost[g.cur] = x
}

func getFreq(s string) []int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[int(s[i]-'a')]++
	}
	return freq
}
