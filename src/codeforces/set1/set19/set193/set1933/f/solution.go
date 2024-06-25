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
		n, m := readTwoNums(reader)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(reader, m)
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const inf = 1 << 30

func solve(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([]int, n)
	fp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}

	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
	}
	dp[0] = 0

	for j := 0; j < m; j++ {
		// 先上下移动
		pq := make(PriorityQueue, n)

		for i := 0; i < n; i++ {
			it := items[i]
			it.index = i
			it.priority = dp[i]
			pq[i] = it

			dp[i] = inf
			fp[i] = inf
		}

		heap.Init(&pq)

		for pq.Len() > 0 && pq[0].priority < inf {
			it := heap.Pop(&pq).(*Item)
			i := it.id
			dp[i] = it.priority
			// go up
			if grid[(i+1)%n][j] == 0 && grid[(i+2)%n][j] == 0 {
				// safe to i+2
				if items[(i+2)%n].priority > it.priority+1 {
					pq.update(items[(i+2)%n], it.priority+1)
				}
			}
		}

		if j+1 == m {
			continue
		}

		// go right
		for i := 0; i < n; i++ {
			if dp[i] < inf && grid[(i+1)%n][j+1] == 0 {
				fp[(i+1)%n] = dp[i] + 1
			}
		}

		copy(dp, fp)
	}

	// 不是dp[n-1]
	res := inf

	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}
		j := dp[i] % n
		// 要算出i当前的位置
		cur := (i - j + n) % n
		diff := n - 1 - cur

		if diff != 0 {
			diff = min(diff, n-diff)
		}

		res = min(res, dp[i]+diff)
	}

	if res == inf {
		return -1
	}

	return res
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
