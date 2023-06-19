package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

func solve(A []int) int {
	n := len(A)
	var best int

	cnt := make(map[int]int)

	for l, r := 0, 0; r < n; r++ {
		cnt[A[r]]++
		for len(cnt) > 2 {
			cnt[A[l]]--
			if cnt[A[l]] == 0 {
				delete(cnt, A[l])
			}
			l++
		}
		best = max(best, r-l+1)
	}

	return best
}

func solve2(A []int) int {
	n := len(A)
	var best int

	newc := 1
	var oldc int
	prev := -1
	cur := A[0]
	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			newc++
		} else if A[i] == prev {
			prev = cur
			cur = A[i]
			oldc += newc
			newc = 1
		} else {
			best = max(best, newc+oldc)
			oldc = newc
			newc = 1
			prev = cur
			cur = A[i]
		}
	}

	return max(best, newc+oldc)
}

func solve1(A []int) int {
	n := len(A)

	mn := NewSegTree(n, 1<<30, min)
	mx := NewSegTree(n, 0, max)

	for i := 0; i < n; i++ {
		mn.Update(i, A[i])
		mx.Update(i, A[i])
	}

	var best int = 1

	for l, r := 0, 1; r < n; r++ {
		a := mn.Get(l, r)
		b := mx.Get(l, r)
		// b - a <= 1
		if A[r] >= a && A[r] <= b || a == b && abs(A[r]-a) == 1 {
			best = max(best, r-l+1)
			continue
		}
		for mx.Get(l, r+1)-mn.Get(l, r+1) > 1 {
			l++
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	return a + b - max(a, b)
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
	seg.arr[p] = seg.op(v, seg.arr[p])
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
