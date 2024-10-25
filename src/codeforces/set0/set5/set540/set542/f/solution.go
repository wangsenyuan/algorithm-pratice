package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	tasks := make([][]int, n)
	for i := 0; i < n; i++ {
		tasks[i] = readNNums(reader, 2)
	}
	res := solve(m, tasks)

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

const inf = 1 << 60

func solve(tot int, tasks [][]int) int {
	n := len(tasks)
	tree := make(PriorityQueue, n)

	for i, task := range tasks {
		depth := tot - task[0]
		value := task[1]
		it := new(Item)
		it.depth = depth
		it.value = value
		it.index = i
		tree[i] = it
	}

	heap.Init(&tree)

	for tree.Len() > 0 {
		first := heap.Pop(&tree).(*Item)
		if first.depth == 0 || tree.Len() == 0 {
			return first.value
		}

		if tree[0].depth == first.depth {
			// merge
			second := heap.Pop(&tree).(*Item)
			first.value += second.value
			first.depth = second.depth - 1
			heap.Push(&tree, first)
		} else {
			first.depth = tree[0].depth
			heap.Push(&tree, first)
		}
	}
	return 0
}

// An Item is something we manage in a priority queue.
type Item struct {
	value int // The value of the item; arbitrary.
	depth int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].depth > pq[j].depth || pq[i].depth == pq[j].depth && pq[i].value > pq[j].value
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

func solve1(tot int, tasks [][]int) int {
	n := len(tasks)

	at := make([][]int, tot+1)

	for _, task := range tasks {
		c, p := task[0], task[1]
		d := tot - c
		at[d] = append(at[d], p)
	}

	for i, vs := range at {
		sort.Ints(vs)
		reverse(vs)
		var ws []int
		var sum int
		for _, v := range vs {
			ws = append(ws, sum)
			sum += v
		}
		ws = append(ws, sum)
		at[i] = ws
	}

	// dp[h][c]表示在第h层，使用n个节点时的最优解
	dp := make([]int, n+1)
	copy(dp, at[tot])

	fp := make([]int, n+1)
	for h := tot - 1; h > 0; h-- {
		for i := 0; i < len(dp); i++ {
			for j := 0; j < len(at[h]); j++ {
				fp[(i+j+1)/2] = max(fp[(i+j+1)/2], dp[i]+at[h][j])
			}
		}
		copy(dp, fp)
		clear(fp)
	}

	if len(at[0]) > 1 {
		return max(dp[1], at[0][1])
	}

	return dp[1]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
