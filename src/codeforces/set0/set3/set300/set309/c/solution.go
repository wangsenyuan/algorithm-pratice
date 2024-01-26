package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	res := solve(a, b)

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

func solve(a []int, b []int) int {
	cnt := make([]int, 30)
	for i := 0; i < len(a); i++ {
		for j := 0; j < 30; j++ {
			if (a[i]>>j)&1 == 1 {
				cnt[j]++
			}
		}
	}

	sort.Ints(b)
	for i := 0; i < len(b); i++ {
		j := b[i]
		for j < 30 && cnt[j] == 0 {
			cnt[j]++
			j++
		}
		if j == 30 {
			return i
		}
		cnt[j]--
		//cnt[b[i]]--
	}

	return len(b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(a []int, b []int) int {
	sort.Ints(b)

	pq := make(IntHeap, 0, len(a))

	check := func(ln int) bool {
		for pq.Len() > 0 {
			heap.Pop(&pq)
		}

		for _, num := range a {
			heap.Push(&pq, num)
		}

		for i := ln - 1; i >= 0; i-- {
			x := 1 << b[i]
			if pq.Len() == 0 || pq[0] < x {
				return false
			}
			y := heap.Pop(&pq).(int)
			y -= x
			if y > 0 {
				heap.Push(&pq, y)
			}
		}

		return true
	}

	l, r := 0, len(b)

	for l < r {
		mid := (l + r + 1) >> 1
		if !check(mid) {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return r
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
