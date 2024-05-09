package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	p := readNNums(reader, n*2)
	pairs := make([][]int, m)
	for i := 0; i < m; i++ {
		pairs[i] = readNNums(reader, 2)
	}

	x := readNum(reader)

	if x == 1 {
		pick := func(u int) int {
			fmt.Printf("%d\n", u)
			return readNum(reader)
		}

		solve1(n, p, pairs, pick)
		return
	}

	var cnt int
	pick := func(u int) int {
		cnt++
		fmt.Printf("%d\n", u)
		if cnt < n {
			return readNum(reader)
		}
		// 作为player 2， 最后一次不需要等待输入
		return 0
	}

	first := readNum(reader)

	solve2(n, p, pairs, first, pick)
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

const oo = 1 << 60

func solve2(n int, p []int, pairs [][]int, first int, pick func(int) int) {
	xs := make([]*Item, len(pairs))
	mandatory := make(Heap, 0, len(pairs))

	opp := make([]int, 2*n+1)
	for i := 0; i <= 2*n; i++ {
		opp[i] = -1
	}

	for i, cur := range pairs {
		a, b := cur[0], cur[1]
		if p[a-1] > p[b-1] {
			a, b = b, a
		}
		xs[i] = new(Item)
		xs[i].id = i
		xs[i].value = p[b-1]
		opp[a] = i
		opp[b] = i
		heap.Push(&mandatory, xs[i])
	}

	items := make([]*Item, 2*n+1)
	nums := make(Heap, 0, 2*n+1)

	for i := 1; i <= 2*n; i++ {
		it := new(Item)
		it.id = i
		it.value = p[i-1]
		items[i] = it
		heap.Push(&nums, it)
	}

	remove := func(u int) {
		if u == 0 || items[u].index < 0 {
			return
		}
		nums.update(items[u], oo)
		heap.Pop(&nums)
	}

	remove(first)

	// 使用的是pairs中的
	for opp[first] >= 0 {
		i := opp[first]
		mandatory.update(xs[i], oo)
		heap.Pop(&mandatory)

		cur := pairs[i]
		a, b := cur[0], cur[1]
		second := first ^ a ^ b
		remove(second)
		first = pick(second)
		remove(first)
	}

	for mandatory.Len() > 0 {
		cur := heap.Pop(&mandatory).(*Item)
		pr := pairs[cur.id]
		u := pr[0]
		if p[pr[0]-1] < p[pr[1]-1] {
			u = pr[1]
		}
		remove(u)
		first = pick(u)
		remove(first)
	}

	for nums.Len() > 0 {
		it := heap.Pop(&nums).(*Item)
		remove(it.id)
		first = pick(it.id)
		remove(first)
	}
}

// play as 1
func solve1(n int, p []int, pairs [][]int, pick func(int) int) {
	use := make([]bool, 2*n+1)

	items := make([]*Item, 2*n+1)

	pq := make(Heap, 0, len(pairs))

	for _, cur := range pairs {
		a, b := cur[0], cur[1]
		if p[a-1] > p[b-1] {
			a, b = b, a
		}
		items[b] = new(Item)
		items[b].id = b
		items[b].value = p[b-1]
		heap.Push(&pq, items[b])
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		v := pick(u)
		// v肯定和u是一对
		use[u] = true
		use[v] = true
	}

	for i := 1; i <= 2*n; i++ {
		if !use[i] {
			it := new(Item)
			it.id = i
			it.value = p[i-1]
			heap.Push(&pq, it)
			items[i] = it
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		v := pick(it.id)
		pq.update(items[v], oo)
		heap.Pop(&pq)
	}
}

type Item struct {
	id    int
	value int
	index int
}

// An IntHeap is a min-heap of ints.
type Heap []*Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Heap) Push(x any) {
	it := x.(*Item)
	it.index = len(*h)
	*h = append(*h, it)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	x.index = -1
	return x
}

func (h *Heap) update(it *Item, v int) {
	it.value = v
	heap.Fix(h, it.index)
}
