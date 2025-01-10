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
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, m, k := readThreeNums(reader)
	friends := make([][]int, m)
	for i := range m {
		friends[i] = readNNums(reader, 2)
	}
	return solve(n, friends, k)
}

func solve(n int, friends [][]int, k int) []int {
	m := len(friends)

	pq := make(PriorityQueue, n)

	items := make([]*Item, n)

	g := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = 0
		items[i].index = i
		pq[i] = items[i]
	}

	for _, cur := range friends {
		u, v := cur[0]-1, cur[1]-1
		items[u].priority++
		items[v].priority++
		g[u][v] = true
		g[v][u] = true
	}

	heap.Init(&pq)

	ans := make([]int, m)

	for i := m - 1; i >= 0; i-- {
		for pq.Len() > 0 && pq[0].priority < k {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			for v := range g[u] {
				if g[v][u] && items[v].index >= 0 {
					delete(g[v], u)
					pq.update(items[v], items[v].priority-1)
				}
			}
		}
		ans[i] = pq.Len()
		cur := friends[i]
		u, v := cur[0]-1, cur[1]-1
		if g[u][v] {
			if items[u].index >= 0 {
				pq.update(items[u], items[u].priority-1)
			}
			delete(g[u], v)
		}
		if g[v][u] {
			if items[v].index >= 0 {
				pq.update(items[v], items[v].priority-1)
			}
			delete(g[v], u)
		}
	}

	return ans
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(v any) {
	it := v.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
