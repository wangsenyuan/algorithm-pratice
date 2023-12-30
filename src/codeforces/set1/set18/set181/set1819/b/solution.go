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
		rects := make([][]int, n)
		for i := 0; i < n; i++ {
			rects[i] = readNNums(reader, 2)
		}
		res := solve(rects)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
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

func solve(rects [][]int) [][]int {
	var area int
	var mw, mh int
	for _, rec := range rects {
		area += rec[0] * rec[1]
		mh = max(mh, rec[0])
		mw = max(mw, rec[1])
	}
	var ans []Pair
	// area = h * w
	if area%mw == 0 {
		ans = append(ans, find(rects, area/mw, mw))
	}
	if area%mh == 0 && area/mh != mw {
		ans = append(ans, find(rects, mh, area/mh))
	}
	var res [][]int
	for i := 0; i < len(ans); i++ {
		if ans[i].first*ans[i].second == area {
			res = append(res, []int{ans[i].first, ans[i].second})
		}
	}
	return res
}

func find(recs [][]int, xh int, xw int) Pair {
	// can we get all the recs from h, w
	n := len(recs)
	hs := make([]int, n)
	ws := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = recs[i][0]
		ws[i] = recs[i][1]
	}
	sh := NewSet(hs)
	sw := NewSet(ws)

	h, w := xh, xw

	for h > 0 && w > 0 {
		if sh.Size() > 0 && sh.First().val == h {
			// could split horizontally
			cur := sh.First()
			sh.Remove(cur.id)
			cur = sw.Remove(cur.id)
			w -= cur.val
		} else if sw.Size() > 0 && sw.First().val == w {
			cur := sw.First()
			sw.Remove(cur.id)
			cur = sh.Remove(cur.id)
			h -= cur.val
		} else {
			// no solution
			return Pair{}
		}
	}

	if h < 0 || w < 0 {
		return Pair{}
	}

	if h == 0 || w == 0 {
		return Pair{xh, xw}
	}
	return Pair{0, 0}
}

type Set struct {
	arr []*Item
	pq  *PriorityQueue
}

func NewSet(val []int) *Set {
	n := len(val)
	arr := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		arr[i] = new(Item)
		arr[i].id = i
		arr[i].val = val[i]
		arr[i].index = i
		pq[i] = arr[i]
	}
	heap.Init(&pq)
	return &Set{arr, &pq}
}

func (set *Set) Size() int {
	return set.pq.Len()
}

func (set *Set) First() *Item {
	return (*(set.pq))[0]
}

func (set *Set) Remove(id int) *Item {
	v := set.arr[id].val
	set.arr[id].val = 1 << 30
	heap.Fix(set.pq, set.arr[id].index)
	heap.Pop(set.pq)
	set.arr[id].val = v
	return set.arr[id]
}

type Item struct {
	id    int
	val   int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val > pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
