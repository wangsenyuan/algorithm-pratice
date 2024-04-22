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

	n, x, y := readThreeNums(reader)

	shows := make([][]int, n)
	for i := 0; i < n; i++ {
		shows[i] = readNNums(reader, 2)
	}

	res := solve(x, y, shows)

	fmt.Println(res)
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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(x int, y int, shows [][]int) int {
	n := len(shows)

	events := make([]Event, 2*n)

	for i, show := range shows {
		events[2*i] = Event{i, show[0], 0}
		events[2*i+1] = Event{i, show[1] + 1, -1}
	}

	slices.SortFunc(events, func(a, b Event) int {
		if a.pos < b.pos || a.pos == b.pos && a.tp < b.tp {
			return -1
		}
		return 1
	})

	tvs := make(PriorityQueue, 0, 1)

	its := make([]*Item, n)

	var res int

	for _, event := range events {
		i := event.id
		pos := event.pos
		tp := event.tp

		if tp == -1 {
			it := its[i]
			it.priority = pos - 1
			heap.Push(&tvs, it)
		} else {
			l, r := shows[i][0], shows[i][1]
			if len(tvs) == 0 || x+(r-l)*y <= (r-tvs[0].priority)*y {
				// res += x + (r - l) * y
				res = add(res, add(x, mul(r-l, y)))
				// just create it
				its[i] = new(Item)
				its[i].id = i
			} else {
				it := heap.Pop(&tvs).(*Item)
				// res += (r - it.priority) * y
				res = add(res, mul(r-it.priority, y))
				its[i] = it
				it.id = i
			}
		}
	}

	return res
}

type Event struct {
	id  int
	pos int
	tp  int
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
