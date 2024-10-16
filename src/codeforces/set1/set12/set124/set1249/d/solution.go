package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}
	res := solve(segs, k)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func solve(segs [][]int, k int) []int {
	n := len(segs)

	pos := make(map[int]int)

	for _, cur := range segs {
		pos[cur[0]]++
		pos[cur[1]+1]++
	}

	arr := make([]int, 0, len(pos))

	for k := range pos {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		pos[arr[i]] = i
	}

	m := len(arr)
	at := make([][]int, m)

	items := make([]*Item, n)

	for i, cur := range segs {
		l, r := cur[0], cur[1]
		l = pos[l]
		r = pos[r+1]

		at[l] = append(at[l], i)
		at[r] = append(at[r], i)

		items[i] = new(Item)
		items[i].id = i
		items[i].priority = r
	}
	occ := make([]int, n)

	active := make(PriorityQueue, 0, n)

	var res []int

	for i := 0; i < m; i++ {
		for _, j := range at[i] {
			occ[j]++
			if occ[j] == 1 {
				heap.Push(&active, items[j])
			} else if items[j].index >= 0 {
				// leave, 有可能已经被删除了
				active.remove(items[j])
			}
		}
		for active.Len() > k {
			it := heap.Pop(&active).(*Item)
			res = append(res, it.id+1)
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
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func (pq *PriorityQueue) remove(it *Item) {
	pq.update(it, 1<<30)
	heap.Pop(pq)
}
