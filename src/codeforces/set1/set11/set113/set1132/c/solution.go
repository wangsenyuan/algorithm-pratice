package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	segs := make([][]int, m)
	for i := range m {
		segs[i] = readNNums(reader, 2)
	}
	return solve(n, segs)
}

type pair struct {
	first  int
	second int
}

func solve(n int, segs [][]int) int {

	cnt := make([]int, n+2)

	for _, seg := range segs {
		l, r := seg[0], seg[1]
		cnt[l]++
		cnt[r+1]--
	}
	v1 := NewSegTree(n + 1)
	v2 := NewSegTree(n + 1)
	var tot int
	for i := 1; i <= n; i++ {
		cnt[i] += cnt[i-1]
		if cnt[i] == 1 {
			v1.Update(i, 1)
		} else if cnt[i] == 2 {
			v2.Update(i, 1)
		}
		if cnt[i] > 0 {
			tot++
		}
	}

	var ans int

	m := len(segs)

	for i := 0; i < m; i++ {
		a, b := segs[i][0], segs[i][1]
		for j := i + 1; j < m; j++ {
			c, d := segs[j][0], segs[j][1]
			if max(a, c) <= min(b, d) {
				// have overlap
				x := min(a, c)
				y := max(a, c)
				tmp := tot
				tmp -= v1.Get(x, y)
				z := min(b, d)
				tmp -= v2.Get(y, z+1)
				w := max(b, d)
				tmp -= v1.Get(z+1, w+1)
				ans = max(ans, tmp)
			} else {
				// no overlap
				tmp := tot
				tmp -= v1.Get(a, b+1)
				tmp -= v1.Get(c, d+1)
				ans = max(ans, tmp)
			}
		}
	}

	return ans
}

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)

	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.arr[p] + seg.arr[p^1]
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	var res int
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res += seg.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += seg.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
