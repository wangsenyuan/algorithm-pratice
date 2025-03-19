package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	p := readNNums(reader, n)
	return solve(p)
}

func solve(p []int) []int {
	n := len(p)
	pos := make([]int, n)
	for i := range n {
		p[i]--
		pos[p[i]] = i
	}

	// 移动距离很好算
	// 就是找到前半段的sum 和 后半段的sum，要找到中间位置
	// 但是 inversion count 要怎么算呢？
	// 新加入一个数，因为它肯定是目前最大的值，所以，只需要知道它后面有多少个比它小就可以了
	sum := NewSegTree(n, 0, func(a, b int) int {
		return a + b
	})

	cnt := NewSegTree(n, 0, func(a, b int) int {
		return a + b
	})

	tr := NewTree(n)

	ans := make([]int, n)

	var inv_count int

	for i := range n {
		j := pos[i]
		// 需要知道j的后面有多少个数比i小
		inv_count += cnt.Get(j, n)
		cnt.Update(j, 1)
		sum.Update(j, j)
		tr.Set(j)

		ans[i] = inv_count

		if i == 0 {
			continue
		}

		l, r := i/2, i/2
		if (i+1)%2 == 0 {
			// 如果是偶数
			r++
		} else {
			// 如果是奇数，最中间的点不用管
			l--
			r++
		}
		m := l + 1
		// l, r 是它们的位置
		l = tr.GetKth(l + 1)
		r = tr.GetKth(r + 1)

		s1 := sum.Get(0, l+1)
		s2 := sum.Get(r, n)
		// 共有这么多项
		ans[i] += s2 - s1
		ans[i] -= m * i
		ans[i] += 2 * m * (m - 1) / 2
	}

	return ans
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

/*
*

	这个是为了找到中间的位置
*/
type Tree struct {
	arr []int
	n   int
}

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	return &Tree{arr, n}
}

func (tr *Tree) Set(p int) {
	// 建节点p设置上去
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.arr[i] = 1
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.arr[i] = tr.arr[2*i+1] + tr.arr[2*i+2]
	}
	loop(0, 0, tr.n-1)
}

func (tr *Tree) GetKth(k int) int {
	var loop func(i int, l int, r int, k int) int

	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if k <= tr.arr[2*i+1] {
			return loop(2*i+1, l, mid, k)
		}
		return loop(2*i+2, mid+1, r, k-tr.arr[2*i+1])
	}

	return loop(0, 0, tr.n-1, k)
}
