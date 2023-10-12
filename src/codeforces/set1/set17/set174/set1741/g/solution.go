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
		f := readNum(reader)
		h := readNNums(reader, f)
		k := readNum(reader)
		p := readNNums(reader, k)

		res := solve(n, E, h, p)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, E [][]int, h []int, p []int) int {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	f := len(h)
	for i := 0; i < f; i++ {
		h[i]--
	}
	k := len(p)
	p = append(p, 0)
	for i := 0; i < k; i++ {
		p[i]--
	}

	masks := make([]int, n)
	for i := 0; i < k; i++ {
		masks[h[p[i]]] |= 1 << i
	}

	sets := make([]map[int]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		sets[i] = make(map[int]bool)
		dist[i] = -1
	}

	sets[0][masks[0]] = true
	dist[0] = 0

	que := make([]int, n)
	var front, tail int
	que[front] = 0
	front++
	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[front] = v
				front++
			}
			// 不能在dist[v] < 0中处理
			if dist[v] == dist[u]+1 {
				for msk := range sets[u] {
					sets[v][msk|masks[v]] = true
				}
			}
		}
	}

	backpack := make([]bool, 1<<k)
	backpack[0] = true
	var j int

	for i := 0; i < f; i++ {
		if i == p[j] {
			// i not have car
			j++
			continue
		}
		nw := make([]bool, 1<<k)
		copy(nw, backpack)
		for msk := 0; msk < 1<<k; msk++ {
			if backpack[msk] {
				for v := range sets[h[i]] {
					nw[msk|v] = true
				}
			}
		}
		backpack = nw
	}
	ans := k

	for mask := 0; mask < 1<<k; mask++ {
		if backpack[mask] {
			var cnt int
			for i := 0; i < k; i++ {
				if mask&(1<<i) == 0 {
					cnt++
				}
			}
			ans = min(ans, cnt)
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
