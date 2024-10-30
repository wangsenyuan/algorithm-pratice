package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		ops := make([][]int, k)
		for i := 0; i < k; i++ {
			ops[i] = readNNums(reader, 3)
		}
		res := solve(n, ops)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, operations [][]int) int {
	m := len(operations)

	slices.SortFunc(operations, func(a, b []int) int {
		return a[0] - b[0]
	})

	items := make([]*Item, m)

	open := make([][]int, n+1)
	end := make([][]int, n+1)
	var dx int
	for i, op := range operations {
		a, d, k := op[0], op[1], op[2]
		open[a] = append(open[a], i)
		end[a+d*k] = append(end[a+d*k], i)
		items[i] = NewItem(i, i)
		dx = max(dx, d)
	}
	set := NewDSU(n + m + 1)

	history := make([][]*PriorityQueue, dx+1)
	for i := 1; i <= dx; i++ {
		history[i] = make([]*PriorityQueue, dx)
		for j := 0; j < dx; j++ {
			tmp := make(PriorityQueue, 0, 1)
			history[i][j] = &tmp
		}
	}
	// 重复的不用管
	for i := 1; i <= n; i++ {

		for _, id := range open[i] {
			op := operations[id]
			a, d := op[0], op[1]
			// 对于后面的位置j来说, 存在 j % d = a % d
			r := a % d

			if history[d][r].Len() > 0 {
				u := (*history[d][r])[0].id
				// 这两个set要union在一起
				set.Union(u+n+1, id+n+1)
			}

			heap.Push(history[d][r], items[id])
		}

		for d := 1; d <= dx; d++ {
			r := i % d
			if history[d][r].Len() > 0 {
				// i = a + ?*d
				// (i - a) % d = 0, i和a同余
				u := (*history[d][r])[0].id
				set.Union(i, u+n+1)
			}
		}

		for _, id := range end[i] {
			// 这个地方需要把id给删除掉，
			op := operations[id]
			a, d := op[0], op[1]
			r := a % d
			history[d][r].remove(items[id])
		}
	}

	res := make(map[int]int)

	for i := 1; i <= n; i++ {
		res[set.Find(i)]++
	}
	return len(res)
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	d := new(DSU)
	d.arr = make([]int, n)
	d.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		d.arr[i] = i
		d.cnt[i]++
	}
	return d
}

func (d *DSU) Find(u int) int {
	if d.arr[u] != u {
		d.arr[u] = d.Find(d.arr[u])
	}
	return d.arr[u]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)
	if a == b {
		return false
	}
	if d.cnt[a] < d.cnt[b] {
		a, b = b, a
	}
	d.cnt[a] += d.cnt[b]
	d.arr[b] = a
	return true
}

func (d *DSU) Reset() {
	for i := 0; i < len(d.arr); i++ {
		d.arr[i] = i
		d.cnt[i] = 1
	}
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

func NewItem(id int, priority int) *Item {
	it := new(Item)
	it.id = id
	it.priority = priority
	return it
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
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) remove(it *Item) {
	pq.update(it, 1<<30)
	heap.Pop(pq)
}
