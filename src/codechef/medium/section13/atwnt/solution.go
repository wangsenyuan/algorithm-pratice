package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	P := readNNums(reader, n-1)

	q := readNum(reader)

	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(n, P, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

const MAX_W = 1000000

func solve(n int, P []int, Q [][]int) []int {
	g := NewGraph(n, n-1)
	cnt := make([]int, n)
	for i := 1; i < n; i++ {
		cnt[P[i-1]-1]++
		g.AddEdge(P[i-1]-1, i)
	}
	var dfs func(p, u int) map[int]int

	mem := make([]map[int]int, n)

	dfs = func(p, u int) map[int]int {
		if cnt[u] == 0 {
			mem[u] = make(map[int]int)
			mem[u][1] = 1
		} else {
			if cnt[u] > 1 {
				mem[u] = make(map[int]int)
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if v == p {
					continue
				}
				tmp := dfs(u, v)

				if cnt[u] == 1 {
					mem[u] = tmp
				} else {
					for k, x := range tmp {
						if int64(k)*int64(cnt[u]) <= MAX_W {
							mem[u][k*cnt[u]] += x
						}
					}
				}
			}
		}

		return mem[u]
	}

	dfs(0, 0)

	ans := make([]int, len(Q))

	for i := 0; i < len(Q); i++ {
		v, w := Q[i][0], Q[i][1]
		v--

		if cnt[v] == 0 {
			// a leaf
			ans[i] = 0
			continue
		}

		var ww int
		for k, x := range mem[v] {
			if w%k == 0 {
				ww += w / k * x
			}
		}
		ans[i] = w - ww
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
	e += 10
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
