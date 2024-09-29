package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	r, _ := os.Open("input.txt")
	reader := bufio.NewReader(r)
	n := readNum(reader)
	buildings := make([][]int32, n)
	for i := 0; i < n; i++ {
		buildings[i] = readNNums(reader, 3)
	}

	res := solve(buildings)

	w, _ := os.Create("output.txt")

	writer := bufio.NewWriter(w)

	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		writer.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
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
	var tmp int
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

func readNNums(reader *bufio.Reader, n int) []int32 {
	res := make([]int32, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		var v int
		x = readInt(bs, x, &v)
		res[i] = int32(v)
	}
	return res
}

type pair struct {
	first  int32
	second int32
	id     int32
}

func solve(buildings [][]int32) [][]int32 {
	n := len(buildings)
	items := make([]*Item, n)
	points := make([]pair, 2*n)
	for i, building := range buildings {
		h, l, r := building[0], building[1], building[2]
		points[2*i] = pair{l, -h, int32(i)}
		points[2*i+1] = pair{r, h, int32(i)}
		it := new(Item)
		it.id = int32(i)
		it.priority = 1 << 30
		items[i] = it
	}

	slices.SortFunc(points, func(a, b pair) int {
		if a.first != b.first {
			return int(a.first) - int(b.first)
		}
		return int(a.second) - int(b.second)
	})

	var res [][]int32

	pq := make(Heap, 0, n+1)
	var prev int32
	heap.Push(&pq, new(Item))

	for _, point := range points {
		x, h, i := point.first, point.second, point.id
		if h < 0 {
			// a left
			items[i].priority = -h
			heap.Push(&pq, items[i])
		} else {
			pq.remove(items[i])
		}

		cur := pq[0].priority

		if cur != prev {
			res = append(res, []int32{x, prev})
			res = append(res, []int32{x, cur})
			prev = cur
		}
	}

	return res
}

type Item struct {
	id       int32
	priority int32
	index    int32
}

type Heap []*Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].priority > h[j].priority }
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = int32(i)
	h[j].index = int32(j)
}

func (h *Heap) Push(x any) {
	it := x.(*Item)
	it.index = int32(len(*h))
	*h = append(*h, it)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) remove(it *Item) {
	it.priority = 1 << 30
	heap.Fix(h, int(it.index))
	heap.Pop(h)
}
