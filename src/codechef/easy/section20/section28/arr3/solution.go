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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		C := readNNums(reader, n)
		k1, k2 := readTwoNums(reader)
		res := solve(A, B, C, k1, k2)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(A, B, C []int, k1, k2 int) int64 {
	n := len(A)
	P := make([]Pair, n)
	var ans int64
	for i := 0; i < n; i++ {
		ans += int64(C[i])
		P[i] = Pair{A[i] - C[i], B[i] - C[i]}
	}
	// order by A[i] - B[i] desc
	sort.Sort(Pairs(P))

	pref := make([]int64, n+1)
	suff := make([]int64, n+1)

	pq := make(IntHeap, 0, n)

	var sum int64

	for i := 0; i < n; i++ {
		if P[i].first > 0 {
			sum += int64(P[i].first)
			heap.Push(&pq, int64(P[i].first))
		}
		if pq.Len() > k1 {
			sum -= heap.Pop(&pq).(int64)
		}
		pref[i+1] = sum
	}

	for pq.Len() > 0 {
		heap.Pop(&pq)
	}

	sum = 0

	for i := n - 1; i >= 0; i-- {
		if P[i].second > 0 {
			sum += int64(P[i].second)
			heap.Push(&pq, int64(P[i].second))
		}
		if pq.Len() > k2 {
			sum -= heap.Pop(&pq).(int64)
		}
		suff[i] = sum
	}

	var add int64

	for i := 0; i <= n; i++ {
		add = max(add, pref[i]+suff[i])
	}

	return ans + add
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first-this[i].second > this[j].first-this[j].second
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// An IntHeap is a min-heap of ints.
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
