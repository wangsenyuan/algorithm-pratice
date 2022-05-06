package main

import (
	"bufio"
	"bytes"
	"container/heap"
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
		n := readNum(reader)
		val := readNNums(reader, n)
		Q := readNum(reader)
		var a, b, c uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &a)
		pos = readUint64(s, pos+1, &b)
		readUint64(s, pos+1, &c)
		M, A, B := solve(n, val, Q, []int64{int64(a), int64(b), int64(c)})
		buf.WriteString(fmt.Sprintf("%d\n%d %d\n", M, A, B))
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

func solve(n int, val []int, Q int, in []int64) (M int64, A int, B int) {
	sort.Ints(val)

	add := make([]int, n)

	for i := 0; i < n; i++ {
		add[i] = val[i] % val[0]
	}

	if val[0] == 1 {
		M = -1
		A = 0
		B = Q
	} else {
		dist := dijkstra(n, val, add)
		for i := 0; i < len(dist); i++ {
			M = max(M, dist[i])
		}
		M -= int64(val[0])

		a, b, c := in[0], in[1], in[2]
		a %= c
		b %= c

		V := int64(val[0])

		st := (a + b) % c
		m1 := st % V
		add1 := a % V
		add2 := (-c%V + V) % V
		for i := 1; i <= Q; i++ {
			if st >= dist[m1] {
				B++
			} else {
				A++
			}
			st = st + a
			m1 = m1 + add1
			if m1 >= V {
				m1 -= V
			}
			if st >= c {
				st -= c
				m1 += add2
				if m1 >= V {
					m1 -= V
				}
			}
		}
	}

	return
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func dijkstra(n int, val []int, add []int) []int64 {
	arr := make([]*Item, val[0])
	pq := make(PQ, val[0])

	for i := 0; i < len(arr); i++ {
		arr[i] = new(Item)
		arr[i].value = i
		arr[i].priority = 1 << 60
		arr[i].index = i
		pq[i] = arr[i]
	}

	arr[0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value
		for i := 0; i < n; i++ {
			v := (add[i] + u) % val[0]
			if arr[v].priority > cur.priority+int64(val[i]) {
				pq.update(arr[v], cur.priority+int64(val[i]))
			}
		}
	}
	dist := make([]int64, val[0])
	for i := 0; i < n; i++ {
		dist[i] = arr[i].priority
	}
	return dist
}

type Item struct {
	priority int64
	value    int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	cur := x.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PQ) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
