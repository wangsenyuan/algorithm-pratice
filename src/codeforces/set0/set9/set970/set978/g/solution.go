package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) ([]int, int, [][]int) {
	n, m := readTwoNums(reader)
	exams := make([][]int, m)
	for i := range m {
		exams[i] = readNNums(reader, 3)
	}
	return solve(n, exams), n, exams
}

func solve(n int, exams [][]int) []int {
	m := len(exams)

	publishAt := make([][]int, n+1)
	eventAt := make([][]int, n+1)

	prepare := make([]int, m)

	for i, exam := range exams {
		s, d, c := exam[0], exam[1], exam[2]
		prepare[i] = c
		publishAt[s] = append(publishAt[s], i)
		eventAt[d] = append(eventAt[d], i)
	}

	active := make(PriorityQueue, 0, n)

	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if len(eventAt[i]) > 1 {
			return nil
		}

		if len(eventAt[i]) == 1 {
			if prepare[eventAt[i][0]] > 0 {
				return nil
			}
			ans[i] = m + 1
		}

		for _, j := range publishAt[i] {
			it := new(Item)
			it.id = j
			it.priority = exams[j][1]
			heap.Push(&active, it)
		}

		if ans[i] == 0 && active.Len() > 0 {
			j := active[0].id
			ans[i] = j + 1
			prepare[j]--
			if prepare[j] == 0 {
				heap.Pop(&active)
			}
		}
	}

	return ans[1:]
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
