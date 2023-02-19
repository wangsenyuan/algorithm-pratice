package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)
		res := solve(s[:n])
		buf.WriteString(fmt.Sprintf("%s\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(s string) string {
	n := len(s)
	pos := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = n
	}
	g := NewGraph(n, 26*n)
	deg := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		for j := 0; j < 26; j++ {
			if pos[j] == n {
				continue
			}
			if j != x-1 && j != x+1 {
				g.AddEdge(i, pos[j])
				deg[pos[j]]++
			}
		}
		pos[x] = i
	}
	pq := make(PQ, 0, n)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			it := new(Item)
			it.id = i
			it.val = s[i]
			heap.Push(&pq, it)
		}
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		buf.WriteByte(cur.val)
		u := cur.id
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				it := new(Item)
				it.id = v
				it.val = s[v]
				heap.Push(&pq, it)
			}
		}
	}

	return buf.String()
}

type Item struct {
	id  int
	val byte
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PQ) Push(x interface{}) {
	i := x.(*Item)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	re := old[n-1]
	*pq = old[:n-1]
	return re
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
