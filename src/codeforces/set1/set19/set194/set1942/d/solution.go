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
		n, k := readTwoNums(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = readNNums(reader, n-i)
		}
		res := solve(n, k, a)
		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, k int, a [][]int) []int {

	A := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		A[i] = make([]int, n+1)
		for j := i; j <= n; j++ {
			A[i][j] = a[i-1][j-i]
		}
	}

	dp := make([][]int, n+1)

	dp[0] = []int{0}

	for i := 1; i <= n; i++ {
		pq := make(PriorityQueue, 0, k)
		heap.Push(&pq, NewItem(i-1, 0, dp[i-1][0]))

		for j := i - 2; j >= -1; j-- {
			var val int
			if j >= 0 {
				val = dp[j][0]
			}
			val += A[j+2][i]
			heap.Push(&pq, NewItem(j, 0, val))
		}
		for pq.Len() > 0 && len(dp[i]) < k {
			it := heap.Pop(&pq).(*Item)
			dp[i] = append(dp[i], it.priority)
			// 最前面那个只能有一次
			if it.id < 0 || it.pos+1 == len(dp[it.id]) {
				continue
			}
			if it.id == i-1 {
				heap.Push(&pq, NewItem(i-1, it.pos+1, dp[i-1][it.pos+1]))
			} else {
				val := dp[it.id][it.pos+1] + A[it.id+2][i]
				heap.Push(&pq, NewItem(it.id, it.pos+1, val))
			}
		}
	}

	return dp[n]
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	pos      int
	priority int
	index    int
}

func NewItem(id int, pos int, priority int) *Item {
	it := new(Item)
	it.id = id
	it.pos = pos
	it.priority = priority
	return it
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
