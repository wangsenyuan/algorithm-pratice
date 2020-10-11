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

	N := readNum(reader)

	ask := func(A, B []int) []int {
		fmt.Printf("1 %d %d\n", len(A), len(B))
		var buf bytes.Buffer
		for i := 0; i < len(A); i++ {
			buf.WriteString(fmt.Sprintf("%d ", A[i]+1))
		}
		buf.WriteByte('\n')
		for i := 0; i < len(B); i++ {
			buf.WriteString(fmt.Sprintf("%d ", B[i]+1))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNNums(reader, 3)
	}

	ans := solve(N, ask)
	fmt.Printf("2 %d\n", ans)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(N int, ask func(A, B []int) []int) int64 {
	arr := make([]int, N)
	cnt := make([]int, N)
	pq := make(PQ, N)
	for i := 0; i < N; i++ {
		arr[i] = i
		cnt[i] = 1
		pq[i] = new(Item)
		pq[i].pos = i
		pq[i].priority = cnt[i]
		pq[i].index = i
	}
	heap.Init(&pq)

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	union := func(a, b int) bool {
		pa := find(a)
		pb := find(b)
		if pa == pb {
			return false
		}
		if cnt[pa] < cnt[pb] {
			pa, pb = pb, pa
		}
		arr[pb] = pa
		cnt[pa] += cnt[pb]
		return true
	}

	var ans int64

	A := make([]int, N)
	B := make([]int, N)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)

		u := cur.pos

		if cnt[find(u)] != cur.priority || cur.priority == N {
			continue
		}
		var x, y int
		for i := 0; i < N; i++ {
			if find(i) == find(u) {
				A[x] = i
				x++
			} else {
				B[y] = i
				y++
			}
		}
		res := ask(A[:x], B[:y])
		a, b, c := res[0], res[1], res[2]
		a--
		b--
		ans += int64(c)
		union(a, b)
		heap.Push(&pq, &Item{pos: find(a), priority: cnt[find(a)]})
	}

	return ans
}

type Item struct {
	pos      int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
