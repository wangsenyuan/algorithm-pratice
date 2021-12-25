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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		T := readNNums(reader, n)
		res := solve(n, A, T)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
		}
	}
	fmt.Print(buf.String())
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

const INF = 1 << 30

func solve(n int, A []int, T []int) [][]int {
	persons := make([]Pair, n)

	for i := 0; i < n; i++ {
		persons[i] = Pair{T[i], i}
	}

	sort.Sort(Persons(persons))

	res := make([][]int, 0, 2*n)
	pq := make(PQ, 0, 2*n)
	R := INF
	for i := 0; i < n; i++ {
		cur := persons[i]
		j := cur.second
		for pq.Len() > 0 && R > T[j] {
			top := heap.Pop(&pq).(*Pair)
			x := min(top.first, R-T[j])
			res = append(res, []int{top.second + 1, R - x, R})
			top.first -= x
			if top.first > 0 {
				heap.Push(&pq, top)
			}
			R -= x
		}
		R = T[j]
		heap.Push(&pq, &Pair{A[j], j})
	}

	for pq.Len() > 0 && R > 0 {
		top := heap.Pop(&pq).(*Pair)
		x := min(top.first, R)
		res = append(res, []int{top.second + 1, R - x, R})
		top.first -= x
		if top.first > 0 {
			heap.Push(&pq, top)
		}
		R -= x
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Persons []Pair

func (this Persons) Len() int {
	return len(this)
}

func (this Persons) Less(i, j int) bool {
	return this[i].first > this[j].first
}

func (this Persons) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type PQ []*Pair

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PQ) Push(x interface{}) {
	p := x.(*Pair)
	*this = append(*this, p)
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	ret := old[n-1]
	*this = old[:n-1]
	return ret
}
