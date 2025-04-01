package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
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

func process(reader *bufio.Reader) (res [][]int, r []int) {
	n := readNum(reader)
	r = readNNums(reader, n)
	res = solve(r)
	return
}

func solve(r []int) [][]int {
	// n := len(r)

	freq := make(map[int]int)
	for _, x := range r {
		freq[x]++
	}

	var pq PQ

	for i, v := range freq {
		it := Item{i, v}
		heap.Push(&pq, it)
	}

	var res [][]int
	for pq.Len() >= 3 {
		a := heap.Pop(&pq).(Item)
		b := heap.Pop(&pq).(Item)
		c := heap.Pop(&pq).(Item)
		delta := c.priority
		if pq.Len() > 0 {
			delta = c.priority - pq[0].priority + 1
		}
		c.priority -= delta
		b.priority -= delta
		a.priority -= delta

		for delta > 0 {
			cur := []int{a.id, b.id, c.id}
			sort.Ints(cur)
			cur[0], cur[2] = cur[2], cur[0]
			res = append(res, cur)
			delta--
		}

		if a.priority > 0 {
			heap.Push(&pq, a)
		}
		if b.priority > 0 {
			heap.Push(&pq, b)
		}
		if c.priority > 0 {
			heap.Push(&pq, c)
		}
	}

	return res
}

type Item struct {
	id       int
	priority int
}

type PQ []Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	it := x.(Item)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	*pq = old[:n-1]
	return old[n-1]
}
