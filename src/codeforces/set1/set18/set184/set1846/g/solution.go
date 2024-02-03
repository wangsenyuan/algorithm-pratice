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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		s := readString(reader)[:n]
		days := make([]int, m)
		medicines := make([][]string, m)
		for i := 0; i < m; i++ {
			days[i] = readNum(reader)
			medicines[i] = make([]string, 2)
			for j := 0; j < 2; j++ {
				medicines[i][j] = readString(reader)[:n]
			}
		}
		res := solve(n, s, days, medicines)
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

const oo = 1 << 60

func solve(n int, s string, days []int, medicines [][]string) int {
	V := 1 << n

	edges := make([][]int, len(medicines))

	for i, cur := range medicines {
		x := convert(cur[0])
		y := convert(cur[1])
		edges[i] = []int{days[i], x ^ (V - 1), y}
	}

	dist := make([]*Item, V)

	pq := make(PriorityQueue, V)

	for i := 0; i < V; i++ {
		dist[i] = new(Item)
		dist[i].id = i
		dist[i].priority = oo
		dist[i].index = i
		pq[i] = dist[i]
	}

	heap.Init(&pq)

	pq.update(dist[convert(s)], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id

		for i := 0; i < len(edges); i++ {
			d, v, w := edges[i][0], edges[i][1], edges[i][2]

			x := (u & v) | w

			if dist[x].priority > it.priority+d {
				pq.update(dist[x], it.priority+d)
			}
		}

	}

	if dist[0].priority == oo {
		return -1
	}
	return dist[0].priority
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func convert(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res <<= 1
		if s[i] == '1' {
			res |= 1
		}
	}
	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, p int) {
	it.priority = p
	heap.Fix(pq, it.index)
}
