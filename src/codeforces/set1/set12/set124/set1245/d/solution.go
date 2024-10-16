package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	cities := make([][]int, n)

	for i := 0; i < n; i++ {
		cities[i] = readNNums(reader, 2)
	}
	c := readNNums(reader, n)
	k := readNNums(reader, n)
	res, stations, edges := solve(cities, c, k)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", res))
	buf.WriteString(fmt.Sprintf("%d\n", len(stations)))
	for _, x := range stations {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteString(fmt.Sprintf("%d\n", len(edges)))
	for _, cur := range edges {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}

	fmt.Print(buf.String())
}
func solve(cities [][]int, c []int, k []int) (res int, stations []int, wires [][]int) {
	n := len(cities)

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dx := abs(cities[i][0] - cities[j][0])
			dy := abs(cities[i][1] - cities[j][1])

			dist[i][j] = (dx + dy) * (k[i] + k[j])
		}
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		items[i] = new(Item)
		items[i].value = i
		items[i].priority = c[i]
		items[i].index = i
		pq[i] = items[i]
	}

	from := make([]int, n)
	for i := 0; i < n; i++ {
		from[i] = -1
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.value
		res += it.priority
		if from[u] < 0 {
			stations = append(stations, u+1)
		} else {
			wires = append(wires, []int{from[u] + 1, u + 1})
		}
		for v := 0; v < n; v++ {
			if items[v].index < 0 {
				continue
			}
			if dist[u][v] < items[v].priority {
				from[v] = u
				pq.update(items[v], dist[u][v])
			}
		}
	}

	return
}

func abs(num int) int {
	return max(num, -num)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
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

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

func (set *DSU) Reset() {
	for i := 0; i < len(set.arr); i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = len(set.arr)
}
