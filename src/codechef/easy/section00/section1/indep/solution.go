package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"sort"
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
		cnt, res := solve(n, m, E)

		buf.WriteString(fmt.Sprintf("%d\n%s\n", cnt, res))
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

func solve(n int, m int, E [][]int) (int, string) {
	g := NewGraph(n, m*2+10)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	r := rand.New(rand.NewSource(99))

	ii := make(map[uint64]int)
	mask := make([]uint64, n)
	for i := 0; i < n; i++ {
		for {
			j := uint64(r.Int63())
			if _, found := ii[j]; !found {
				ii[j] = i
				mask[i] = j
				break
			}
		}
	}

	groups := make(map[uint64]int)
	last := make(map[uint64]int)
	count := make(map[uint64]int)
	que := make([]int, n)
	for i := 0; i < n; i++ {
		var end int
		for j := g.nodes[i]; j > 0; j = g.next[j] {
			v := g.to[j]
			que[end] = v
			end++
		}
		sort.Ints(que[:end])
		var tmp uint64

		for j := 0; j < end; j++ {
			tmp = tmp*19 + mask[que[j]]
		}

		count[tmp] = end
		groups[tmp]++
		last[tmp] = i
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '1'
	}
	var res int
	for k, v := range count {
		if groups[k]+v == n {
			res++
			if res == 1 {
				u := last[k]
				for j := g.nodes[u]; j > 0; j = g.next[j] {
					buf[g.to[j]] = '0'
				}
			}
		}
	}

	if res == 0 {
		for i := 0; i < n; i++ {
			buf[i] = '0'
		}
	}

	return res, string(buf)
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
