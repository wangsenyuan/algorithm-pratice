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
		n := readNum(reader)
		a := readNNums(reader, n)
		ok, res := solve(a)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
			}
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int) (bool, [][]int) {
	//  如果 n = 2
	// a[0] < a[1]
	// 1, 2 => false
	// 2, 3 => 2, 2
	n := len(a)
	ok := true
	for i := 0; i < n; i++ {
		if a[i] != a[0] {
			ok = false
			break
		}
	}
	if ok {
		return true, nil
	}

	for i := 0; i < n; i++ {
		if a[i] == 1 {
			return false, nil
		}
	}
	var res [][]int

	pq := make(PriorityQueue, n)
	x := &Item{0, a[0]}
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.value = a[i]
		if a[i] < x.value {
			x.id = i
			x.value = a[i]
		}
		pq[i] = it
	}

	heap.Init(&pq)

	for pq.Len() > 0 && pq[0].value > x.value {
		cur := heap.Pop(&pq).(*Item)
		res = append(res, []int{cur.id + 1, x.id + 1})

		cur.value = (cur.value + x.value - 1) / x.value
		if cur.value < x.value {
			x = cur
		}
		heap.Push(&pq, cur)
	}

	return true, res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int
	value int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
