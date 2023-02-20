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
	n, m := readTwoNums(reader)

	res := solve(n, m)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(n int, m int) []int {

	res := make([]int, n)
	var i int
	lf := (m + 1) / 2
	rg := lf
	var expect int
	if m&1 == 0 {
		rg++
	} else {
		res[0] = lf
		lf--
		rg++
		i = 1
		expect = 1
	}

	for i < m && i < n {
		if i&1 == expect {
			res[i] = lf
			lf--
		} else {
			res[i] = rg
			rg++
		}

		i++
	}

	for i < n {
		res[i] = res[i-m]
		i++
	}

	return res
}

func solve1(n int, m int) []int {
	// n balls, m baskets
	mid := (m + 1) / 2

	pq := make(PriorityQueue, m)
	for i := 1; i <= m; i++ {
		basket := new(Item)
		basket.id = i
		basket.cnt = 0
		if m&1 == 1 {
			basket.pos = abs(mid - i)
		} else {
			if i <= mid {
				basket.pos = mid - i
			} else {
				basket.pos = i - (mid + 1)
			}
		}
		basket.index = i - 1
		pq[i-1] = basket
	}
	heap.Init(&pq)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = pq[0].id
		pq.update(pq[0], pq[0].cnt+1)
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Item struct {
	id    int
	cnt   int
	pos   int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cnt < pq[j].cnt || pq[i].cnt == pq[j].cnt && pq[i].pos < pq[j].pos || pq[i].cnt == pq[j].cnt && pq[i].pos == pq[j].pos && pq[i].id < pq[j].id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(it *Item, cnt int) {
	it.cnt = cnt
	heap.Fix(pq, it.index)
}
