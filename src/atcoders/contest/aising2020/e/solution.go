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
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	camels := make([][]int, n)
	for i := 0; i < n; i++ {
		camels[i] = readNNums(reader, 3)
	}
	return solve(n, camels)
}

func solve(n int, camels [][]int) int {
	lf := make([][]int, n+1)
	rt := make([][]int, n+2)
	var res int
	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		k := camels[i][0]
		l, r := camels[i][1], camels[i][2]
		if l > r {
			lf[k] = append(lf[k], i)
			cnt[0]++
			res += r
		} else if l < r {
			// 这些希望排到右边去
			rt[k+1] = append(rt[k+1], i)
			cnt[1]++
			res += l
		} else {
			res += l
			cnt[2]++
		}
	}
	// cnt[0]个排到左边，cnt[1]个排到右边
	// cnt[0] 和 cnt[1] 不会重叠的
	active := make(PriorityQueue, 0, cnt[0])
	for k := n; k > 0; k-- {
		for _, i := range lf[k] {
			it := new(Item)
			it.id = i
			it.priority = camels[i][1] - camels[i][2]
			heap.Push(&active, it)
		}
		if k <= cnt[0] && active.Len() > 0 {
			it := heap.Pop(&active).(*Item)
			res += it.priority
		}
	}
	for active.Len() > 0 {
		heap.Pop(&active)
	}
	for k := 1; k <= n; k++ {
		for _, i := range rt[k] {
			it := new(Item)
			it.id = i
			it.priority = camels[i][2] - camels[i][1]
			heap.Push(&active, it)
		}
		if k >= n-cnt[1] && active.Len() > 0 {
			it := heap.Pop(&active).(*Item)
			res += it.priority
		}
	}

	return res
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
	return pq[i].priority > pq[j].priority
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
func (pq *PriorityQueue) update(item *Item, p int) {
	item.priority = p
	heap.Fix(pq, item.index)
}
