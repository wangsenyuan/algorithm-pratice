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

	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	res := solve(n, E)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}

	fmt.Println(buf.String())
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, E [][]int) []int64 {

	dist := make([]int64, n)

	var sum int64
	var edges [][]int
	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		if u == 0 || v == 0 {
			sum += int64(w)
			dist[u^v] = int64(w)
		} else {
			edges = append(edges, []int{u, v, w})
		}
	}

	todo := make(PriorityQueue, 0, len(edges))
	set := NewDSU(n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		val := int64(w) - max(dist[set.Find(u)], dist[set.Find(v)])
		it := new(Item)
		it.u = u
		it.v = v
		it.w = w
		it.priority = val
		heap.Push(&todo, it)
	}
	ans := make([]int64, n-1)
	ans[n-2] = sum

	merge := func(u, v int) {
		u = set.Find(u)
		v = set.Find(v)
		set.Union(u, v)
		x := min(dist[u], dist[v])
		dist[u] = x
		dist[v] = x
	}

	for i := n - 3; i >= 0; {
		cur := heap.Pop(&todo).(*Item)
		u, v, w := cur.u, cur.v, cur.w
		if set.Find(u) == set.Find(v) {
			continue
		}

		val := int64(w) - max(dist[set.Find(u)], dist[set.Find(v)])
		if cur.priority != val {
			cur.priority = val
			heap.Push(&todo, cur)
			continue
		}
		sum += val
		ans[i] = sum
		merge(u, v)
		i--
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	u        int
	v        int
	w        int
	priority int64
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type DSU struct {
	arr  []int
	cnt  []int
	find func(int) int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	return &DSU{arr, cnt, find}
}

func (dsu *DSU) Find(x int) int {
	return dsu.find(x)
}

func (dsu *DSU) Union(x, y int) bool {
	x = dsu.find(x)
	y = dsu.find(y)
	if x == y {
		return false
	}
	if dsu.cnt[x] < dsu.cnt[y] {
		x, y = y, x
	}
	dsu.cnt[x] += dsu.cnt[y]
	dsu.arr[y] = x
	return true
}
