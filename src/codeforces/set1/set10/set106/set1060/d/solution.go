package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	guests := make([][]int, n)
	for i := 0; i < n; i++ {
		guests[i] = readNNums(reader, 2)
	}

	res := solve(guests)

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

func solve(guests [][]int) int {
	n := len(guests)
	// 至少需要n个椅子
	res := n

	it1 := make([]*Item, n)
	it2 := make([]*Item, n)
	pq1 := make(PriorityQueue, n)
	pq2 := make(PriorityQueue, n)

	add := func(i int, v int, its []*Item, pq PriorityQueue) {
		it := new(Item)
		it.id = i
		it.value = v
		it.index = i
		its[i] = it
		pq[i] = it
	}

	for i := 0; i < n; i++ {
		guest := guests[i]
		l, r := guest[0], guest[1]
		add(i, r, it1, pq1)
		add(i, l, it2, pq2)
	}

	heap.Init(&pq1)
	heap.Init(&pq2)
	// len(pq1) == len(pq2)
	for pq1.Len() > 0 {
		a := heap.Pop(&pq1).(*Item)
		b := heap.Pop(&pq2).(*Item)
		res += max(a.value, b.value)
		if a.id == b.id {
			continue
		}
		nl := guests[a.id][0]
		nr := guests[b.id][1]
		pq1.update(a, nr)
		pq2.update(b, nl)
	}

	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int
	value int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value > pq[j].value
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
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}
