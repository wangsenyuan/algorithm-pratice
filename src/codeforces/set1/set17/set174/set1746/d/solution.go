package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		P := readNNums(reader, n-1)
		S := readNNums(reader, n)
		res := solve(n, k, P, S)
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

func solve(n, k int, P []int, S []int) int64 {
	g := NewGraph(n, 2*n)

	deg := make([]int, n)
	for i := 0; i < n-1; i++ {
		p := P[i] - 1
		g.AddEdge(p, i+1)
		g.AddEdge(i+1, p)
		deg[i+1]++
		deg[p]++
	}

	dp := make([][]Pair, n)

	var dfs2 func(p int, u int, k int) int64

	dfs2 = func(p int, u int, k int) int64 {
		if k == 0 {
			return 0
		}
		for _, cur := range dp[u] {
			if cur.first == k {
				return cur.second
			}
		}
		tmp := int64(k) * int64(S[u])
		cnt := deg[u]
		if p >= 0 {
			cnt--
		}
		if cnt == 0 {
			return tmp
		}
		x := k / cnt
		if k%cnt == 0 {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					tmp += dfs2(u, v, x)
				}
			}
		} else {
			dp1 := make([]int64, 0, cnt)
			dp2 := make([]int64, 0, cnt)
			id := make([]int, 0, cnt)
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					dp1 = append(dp1, dfs2(u, v, x))
					dp2 = append(dp2, dfs2(u, v, x+1))
					id = append(id, len(id))
				}
			}

			sort.Slice(id, func(i, j int) bool {
				a := dp2[id[i]] - dp1[id[i]]
				b := dp2[id[j]] - dp1[id[j]]
				return a > b
			})
			r := k % cnt
			for i := 0; i < cnt; i++ {
				if i < r {
					tmp += dp2[id[i]]
				} else {
					tmp += dp1[id[i]]
				}
			}
		}

		dp[u] = append(dp[u], Pair{k, tmp})

		return tmp
	}

	return dfs2(-1, 0, k)
}

type Pair struct {
	first  int
	second int64
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
