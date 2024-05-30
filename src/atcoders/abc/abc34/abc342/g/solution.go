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

	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			queries[i] = make([]int, 4)
			queries[i][0] = 1
		} else if s[0] == '2' {
			queries[i] = make([]int, 2)
			queries[i][0] = 2
		} else {
			queries[i] = make([]int, 2)
			queries[i][0] = 3
		}
		pos := 2
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	res := solve(a, queries)

	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(a []int, queries [][]int) []int {
	type PQ []IntHeap

	n := len(a)

	arr := make([]PQ, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = make(PQ, 2)
		for j := 0; j < 2; j++ {
			arr[i][j] = make(IntHeap, 0, 1)
		}
	}

	add := func(x int) func(pq PQ) {
		return func(pq PQ) {
			heap.Push(&pq[0], x)
		}
	}

	cancel := func(x int) func(pq PQ) {
		return func(pq PQ) {
			heap.Push(&pq[1], x)
		}
	}

	update := func(l int, r int, f func(pq PQ)) {
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				f(arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				f(arr[r])
			}
			l >>= 1
			r >>= 1
		}
	}

	top_or_zero := func(pq PQ) int {
		for pq[0].Len() > 0 && pq[1].Len() > 0 && pq[0][0] == pq[1][0] {
			heap.Pop(&pq[0])
			heap.Pop(&pq[1])
		}

		if pq[0].Len() == 0 {
			return 0
		}
		return pq[0][0]
	}

	get := func(i int) int {
		i += n
		var res int
		for i > 0 {
			res = max(res, top_or_zero(arr[i]))
			i >>= 1
		}
		return res
	}

	for i := 0; i < n; i++ {
		update(i, i+1, add(a[i]))
	}

	var res []int

	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			l--
			update(l, r, add(x))
		} else if cur[0] == 2 {
			i := cur[1]
			qr := queries[i-1]
			l, r, x := qr[1], qr[2], qr[3]
			l--
			update(l, r, cancel(x))
		} else {
			i := cur[1]
			res = append(res, get(i-1))
		}
	}

	return res

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}