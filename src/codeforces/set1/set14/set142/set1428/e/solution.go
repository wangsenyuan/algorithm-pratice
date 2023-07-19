package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(k, A)
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
func solve(k int, A []int) int64 {
	n := len(A)
	pq := make(PQ, n)

	for i := 0; i < n; i++ {
		pq[i] = Carrot{A[i], 1}
	}

	heap.Init(&pq)

	cnt := n
	for cnt < k {
		cur := heap.Pop(&pq).(Carrot)
		cnt++
		cur = cur.CutMore()
		heap.Push(&pq, cur)
	}

	var res int64

	for i := 0; i < len(pq); i++ {
		cur := pq[i]
		res += calc(cur.ln, cur.cuts)
	}

	return res
}

type Carrot struct {
	ln   int
	cuts int
}

func calc(ln int, cuts int) int64 {
	// 一共进行cuts切割，使得它们尽可能的相等
	x := ln / cuts
	if ln%cuts == 0 {
		return int64(cuts) * int64(x) * int64(x)
	}
	// 有一部分是x长度的,有一部分是 x + 1长度的
	y := ln - x*cuts
	res := int64(y) * int64(x+1) * int64(x+1)
	res += int64(cuts-y) * int64(x) * int64(x)
	return res
}

func (this Carrot) GainWhenCutMore() int64 {
	return calc(this.ln, this.cuts) - calc(this.ln, this.cuts+1)
}

func (this Carrot) Less(that Carrot) bool {
	return this.GainWhenCutMore() > that.GainWhenCutMore()
}

func (this Carrot) CutMore() Carrot {
	return Carrot{this.ln, this.cuts + 1}
}

// An IntHeap is a min-heap of ints.
type PQ []Carrot

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Carrot))
}

func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
