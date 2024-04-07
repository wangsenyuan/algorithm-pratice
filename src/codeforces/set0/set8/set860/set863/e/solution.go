package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	tvs := make([][]int, n)

	for i := 0; i < n; i++ {
		tvs[i] = readNNums(reader, 2)
	}

	res := solve(tvs)

	fmt.Println(res)
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

func solve(tvs [][]int) int {
	id := make(map[int]int)
	for _, tv := range tvs {
		id[tv[0]]++
		//id[tv[1]]++
		id[tv[1]+1]++
	}
	arr := make([]int, 0, len(id))
	for k := range id {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	n := len(arr)
	for i := 0; i < n; i++ {
		id[arr[i]] = i
	}

	diff := make([]int, n+1)

	for _, tv := range tvs {
		l, r := tv[0], tv[1]+1
		diff[id[l]]++
		diff[id[r]]--
	}

	set := NewSegTree(n, n, min)
	set.Update(0, diff[0])
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
		set.Update(i, diff[i])
	}

	for i, tv := range tvs {
		l, r := tv[0], tv[1]+1
		l = id[l]
		r = id[r]
		if set.Get(l, r) >= 2 {
			return i + 1
		}
	}

	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
