package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) []string {
	n := readNum(reader)
	s := make([]string, n)
	for i := range n {
		s[i] = readString(reader)
	}
	sep := readString(reader)
	return solve(n, s, sep)
}

type Item struct {
	id    int
	val   string
	index int
}

func solve(n int, s []string, sep string) []string {
	items := make([]*Item, n)
	// by length
	trs := make([]*PQ, 11)
	for i := 1; i < 11; i++ {
		tr := make(PQ, 0, n)
		trs[i] = &tr
	}
	var tot int
	for i := range n {
		x := s[i]
		it := new(Item)
		it.id = i
		it.val = x
		heap.Push(trs[len(x)], it)
		items[i] = it
		tot += len(x)
	}
	// 每行的长度
	tot /= (n / 2)

	slices.SortFunc(items, func(a, b *Item) int {
		// 看看哪个应该在前面
		x := a.val + sep + b.val
		y := b.val + sep + a.val
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	})

	var res []string

	marked := make([]bool, n)

	for _, cur := range items {
		if marked[cur.id] {
			// already used
			continue
		}
		x := cur.val
		// 先把x删除掉
		trs[len(x)].remove(cur)
		ly := tot - len(x)
		//  如果能有答案，那么y肯定是存在的
		it := heap.Pop(trs[ly]).(*Item)
		y := it.val
		res = append(res, x+sep+y)
		marked[it.id] = true
	}

	return res
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	// val是相同长度的
	return pq[i].val < pq[j].val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) remove(it *Item) {
	it.val = ""
	heap.Fix(pq, it.index)
	heap.Pop(pq)
}
