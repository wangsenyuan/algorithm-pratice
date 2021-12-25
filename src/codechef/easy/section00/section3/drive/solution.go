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
		n, m := readTwoNums(reader)
		W := make([][]int, m)
		for i := 0; i < m; i++ {
			W[i] = readNNums(reader, 2)
		}
		res := solve(n, m, W)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(N, M int, workers [][]int) int {
	segs := make([]Pair, M)

	for i, cur := range workers {
		segs[i] = Pair{cur[0], cur[1]}
	}

	sort.Sort(Pairs(segs))

	pq := make(PQ, 0, M)

	check := func(k int) bool {
		// each worker cleans at most k-dist
		cur := 1
		var i int
		for cur < N {
			for i < len(segs) && segs[i].first <= cur {
				heap.Push(&pq, NewItem(i, segs[i].second))
				i++
			}
			tmp := cur
			for tmp == cur && pq.Len() > 0 {
				next := heap.Pop(&pq).(*Item).prior
				cur = max(cur, min(next, cur+k))
			}
			if cur == tmp {
				break
			}
		}

		for pq.Len() > 0 {
			heap.Pop(&pq)
		}

		return cur == N
	}

	if !check(N) {
		return -1
	}

	var left, right = 1, N
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Item struct {
	value int
	prior int
	index int
}

func NewItem(v, p int) *Item {
	item := new(Item)
	item.value = v
	item.prior = p
	return item
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].prior < pq[j].prior
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(item interface{}) {
	x := item.(*Item)
	x.index = len(*pq)
	*pq = append(*pq, x)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	res := old[len(old)-1]
	res.index = -1
	*pq = old[:len(old)-1]
	return res
}
