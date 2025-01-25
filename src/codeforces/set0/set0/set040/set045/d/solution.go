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
	_, res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (tasks [][]int, res []int) {
	n := readNum(reader)
	tasks = make([][]int, n)
	for i := range n {
		tasks[i] = readNNums(reader, 2)
	}

	res = solve(tasks)

	return
}

func solve(tasks [][]int) []int {
	n := len(tasks)
	var x int
	for _, task := range tasks {
		x = max(x, task[1])
	}
	x++
	open := make([][]int, x)

	for i, task := range tasks {
		l := task[0]
		open[l] = append(open[l], i)
	}
	pq := make(PriorityQueue, 0, n)
	ans := make([]int, n)
	var cnt int
	for day := 1; day < x; day++ {
		for _, i := range open[day] {
			it := new(Item)
			it.id = i
			it.priority = tasks[i][1]
			heap.Push(&pq, it)
		}
		if pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			ans[it.id] = day
			cnt++
		}
		if cnt == n {
			break
		}
	}
	return ans
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
	return pq[i].priority < pq[j].priority
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
