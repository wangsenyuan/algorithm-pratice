package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	bridges := make([][]int, m)
	for i := range m {
		bridges[i] = readNNums(reader, 2)
	}
	return solve(n, bridges)
}

func solve(n int, bridges [][]int) string {
	m := len(bridges)
	g := NewGraph(n, n+m)
	next := make([]int, n)
	for i := 0; i+1 < n; i++ {
		g.AddEdge(i, i+1)
		next[i] = i + 1
	}
	next[n-1] = n
	for _, bridge := range bridges {
		u, v := bridge[0]-1, bridge[1]-1
		g.AddEdge(u, v)
		next[u] = max(next[u], v)
	}

	dist := make([]int, n)
	for i := range n {
		dist[i] = -1
	}
	dist[0] = 0
	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v < n && dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	cnt := NewSegTree(n)

	for u := 0; u+1 < n; u++ {
		v := next[u]
		// 这里不能用dist[v], 而要用dist[u] + 1
		// 因为 dist[v] 可能 < dist[u] + 1, 且这里的目的是找到，到达节点u以后，哪些开始的节点s是安全的(对于ellise)
		w := v - (dist[u] + 1)
		if u+1 < w {
			// 当s处在区间 u+1, u+2....w-1 上时，ellise 可以更早的通过节点u到达节点v, 从而获胜
			// 但不包括节点u
			cnt.Update(u+1, w, 1)
		}
	}

	ans := make([]byte, n-1)

	for u := range n - 1 {
		if cnt.Get(u) > 0 {
			ans[u] = '0'
		} else {
			ans[u] = '1'
		}
	}

	return string(ans)
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

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, n*2)
	return &SegTree{arr, n}
}

func (tr *SegTree) Update(l int, r int, v int) {
	l += tr.sz
	r += tr.sz
	for l < r {
		if l&1 == 1 {
			tr.arr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr.arr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr *SegTree) Get(p int) int {
	p += tr.sz
	var res int
	for p > 0 {
		res += tr.arr[p]
		p >>= 1
	}
	return res
}
