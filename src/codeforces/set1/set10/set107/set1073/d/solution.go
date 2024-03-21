package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, t := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, t)

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

func solve(a []int, tot int) int {
	n := len(a)
	sum := NewSegTree(n)
	cnt := NewSegTree(n)

	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		sum.Update(i, a[i])
		cnt.Update(i, 1)
		it := new(Item)
		it.id = i
		it.priority = a[i]
		pq[i] = it
	}

	heap.Init(&pq)

	cur := tot
	var res int
	pos := 0

	for cur > 0 && pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		i := it.id

		if it.priority > cur {
			// no use
			sum.Update(i, 0)
			cnt.Update(i, 0)
			continue
		}
		if pos < i {
			tmp := sum.Get(pos, i)

			if tmp > cur {
				// not able to reach here
				// cancel it
				sum.Update(i, 0)
				cnt.Update(i, 0)
				continue
			}
			res += cnt.Get(pos, i)
			cur -= tmp
		} else if pos > i {
			tmp := sum.Get(pos, n) + sum.Get(0, i)
			if tmp > cur {
				sum.Update(i, 0)
				cnt.Update(i, 0)
				continue
			}
			res += cnt.Get(pos, n) + cnt.Get(0, i)
			cur -= tmp
		}
		all := sum.Get(0, n)
		// 如果从位置i开始, 循环一遍至少减少这么多
		x := cur / all
		cur -= all * x
		res += x * cnt.Get(0, n)
		if cur >= it.priority {
			// 它还可以被用一次
			res++
			cur -= it.priority
		}
		sum.Update(i, 0)
		cnt.Update(i, 0)
		pos = i
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id < pq[j].id
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
