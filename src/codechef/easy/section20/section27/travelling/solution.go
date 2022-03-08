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
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, E [][]int) int {
	g := NewGraph(n, len(E)*2+1)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, 2*n)
	front, end := n, n
	que[end] = 0
	end++
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[0] = 0
	for front < end {
		u := que[front]
		front++
		if u == n-1 {
			break
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 || dist[u] < dist[v] {
				dist[v] = dist[u]
				front--
				que[front] = v
			}
		}
		if u-1 > 0 && dist[u-1] < 0 {
			dist[u-1] = dist[u] + 1
			que[end] = u - 1
			end++
		}
		if u+1 < n && dist[u+1] < 0 {
			dist[u+1] = dist[u] + 1
			que[end] = u + 1
			end++
		}
	}

	return dist[n-1]
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

// 这个算法是错的， 考虑一开始的集合 {1, 4}, {2,5}, {3, 6}
// 用这个算法得到的结果是3， 答案应该是2
func solve1(n int, E [][]int) int {

	arr := make([]int, n)
	cnt := make([]int, n)
	right_most := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
		right_most[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	union := func(a, b int) bool {
		pa := find(a)
		pb := find(b)
		if pa == pb {
			return false
		}
		if cnt[pa] < cnt[pb] {
			pa, pb = pb, pa
		}
		cnt[pa] += cnt[pb]
		arr[pb] = pa
		right_most[pa] = max(right_most[pa], right_most[pb])
		return true
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		union(u, v)
	}
	var ans int
	for i := right_most[find(0)]; i < n-1; i = right_most[find(i)] {
		union(i, i+1)
		if i > 0 {
			union(i, i-1)
		}
		ans++
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
