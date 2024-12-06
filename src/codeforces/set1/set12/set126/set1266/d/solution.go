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
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) [][]int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) [][]int {
	bal := make([]int, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		bal[u] += w
		bal[v] -= w
	}

	out := make(PriorityQueue, 0, n)
	in := make(PriorityQueue, 0, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.priority = bal[i]
		it.index = i
		if bal[i] > 0 {
			heap.Push(&out, it)
		} else if bal[i] < 0 {
			heap.Push(&in, it)
		}
	}

	var res [][]int

	for in.Len() > 0 && out.Len() > 0 {
		a := heap.Pop(&in).(*Item)
		b := heap.Pop(&out).(*Item)
		val := min(b.priority, abs(a.priority))
		res = append(res, []int{b.id + 1, a.id + 1, val})
		a.priority += val
		b.priority -= val
		if a.priority != 0 {
			heap.Push(&in, a)
		}
		if b.priority != 0 {
			heap.Push(&out, b)
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
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
	return abs(pq[i].priority) > abs(pq[j].priority)
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
