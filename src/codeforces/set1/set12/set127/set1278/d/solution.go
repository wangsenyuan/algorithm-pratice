package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) bool {
	n := readNum(reader)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = readNNums(reader, 2)
	}
	return solve(segs)
}

func solve(segs [][]int) bool {
	// 左边的好算，右边的也好算
	// 中间的怎么算呢？
	n := len(segs)

	slices.SortFunc(segs, func(a, b []int) int {
		return a[0] - b[0]
	})

	items := make([]*Item, n)
	end := make([][]int, n*2+1)

	for i, cur := range segs {
		it := new(Item)
		it.id = i
		it.priority = cur[1]
		it.index = -1
		items[i] = it
		end[cur[1]] = append(end[cur[1]], i)
	}

	set := NewDSU(n)

	active := make(PriorityQueue, 0, n)
	var pos int
	for u, cur := range segs {
		l, r := cur[0], cur[1]

		for pos < l {
			for _, i := range end[pos] {
				active.remove(items[i])
			}
			pos++
		}

		// 在区间r-1范围内的，都是可以连起来的
		var arr []*Item
		for active.Len() > 0 && active[0].priority < r {
			it := heap.Pop(&active).(*Item)
			if !set.Union(u, it.id) {
				return false
			}
			arr = append(arr, it)
		}
		for _, it := range arr {
			heap.Push(&active, it)
		}
		heap.Push(&active, items[u])
	}

	return set.size == 1
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
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

func (pq *PriorityQueue) remove(it *Item) {
	old := it.priority
	it.priority = -1
	heap.Fix(pq, it.index)
	heap.Pop(pq)
	it.priority = old
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
