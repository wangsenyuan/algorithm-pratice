package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--

		n, m, k := readThreeNums(reader)
		C := readNNums(reader, n)
		T := readNNums(reader, n)
		res := solve(n, m, k, T, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n, m, K int, T []int, C []int) int {
	cnt := make([]int, m)
	for i := 0; i < n; i++ {
		cnt[C[i]-1]++
	}
	ii := make([][]int, m)

	for i := 0; i < m; i++ {
		ii[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		j := C[i] - 1
		ii[j] = append(ii[j], T[i])
	}

	for i := 0; i < m; i++ {
		sort.Ints(ii[i])
	}

	P := make([]int, m)

	subjects := make([]*Item, m)
	pq := make(PriorityQueue, m)

	for i := 0; i < m; i++ {
		cur := new(Item)
		cur.pos = i
		cur.priority = math.MaxInt32
		if len(ii[i]) > 0 {
			cur.priority = ii[i][0]
			P[i]++
		}
		cur.index = i
		subjects[i] = cur
		pq[i] = cur
	}

	heap.Init(&pq)

	var time int
	var score int
	for time < K && pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		if cur.priority == math.MaxInt32 {
			break
		}
		if time+cur.priority <= K {
			score++
			time += cur.priority
		}
		j := cur.pos
		if P[j]+2 <= len(ii[j]) {
			cur.priority = ii[j][P[j]] + ii[j][P[j]+1]
			P[j] += 2
			heap.Push(&pq, cur)
		}
	}

	return score
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos      int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
