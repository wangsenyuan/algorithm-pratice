package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	x := readNNums(reader, m)
	res := solve(n, x)
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

func solve(n int, x []int) []int {
	ans := make([]int, n+2)

	m := len(x)

	for i := 0; i+1 < m; i++ {
		a, b := x[i], x[i+1]
		if a == b {
			continue
		}
		u, v := min(a, b), max(a, b)
		// a, b 也是它们的pos
		// 如果是u前面的数，移动到了首位，那么v-u的贡献就是 v - u
		ans[1] += v - u
		ans[u] -= v - u
		// 如果是v后的数移动到了首位，贡献也是 v - u
		ans[v+1] += v - u
		// 如果u移动到第一位
		ans[u] += v - 1
		ans[u+1] -= v - 1
		// 如果v移动到第一位
		ans[v] += u
		ans[v+1] -= u
		ans[u+1] += v - u - 1
		ans[v] -= v - u - 1
	}
	for i := 1; i <= n; i++ {
		ans[i] += ans[i-1]
	}
	return ans[1 : n+1]
}

func solve2(n int, x []int) []int {

	tr := NewSegTree2(n + 1)

	m := len(x)

	for i := 0; i+1 < m; i++ {
		a, b := x[i], x[i+1]
		if a == b {
			continue
		}
		u, v := min(a, b), max(a, b)
		// a, b 也是它们的pos
		if u > 1 {
			// 如果是u前面的数，移动到了首位，那么v-u的贡献就是 v - u
			tr.Update(1, u, v-u)
		}
		if v < n {
			// 如果是v后的数移动到了首位，贡献也是 v - u
			tr.Update(v+1, n+1, v-u)
		}

		// 如果u移动到第一位
		tr.Update(u, u+1, v-1)
		// 如果v移动到第一位
		tr.Update(v, v+1, u)

		if u+1 < v {
			tr.Update(u+1, v, v-u-1)
		}
	}
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = tr.Get(i)
	}
	return ans[1:]
}

func abs(num int) int {
	return max(num, -num)
}

type SegTree2 struct {
	arr []int
	sz  int
}

func NewSegTree2(n int) *SegTree2 {
	tr := new(SegTree2)
	tr.arr = make([]int, 2*n)
	tr.sz = n
	return tr
}

func (tr *SegTree2) Update(l int, r int, v int) {
	l += tr.sz
	r += tr.sz

	for l < r {
		if l&1 == 1 {
			tr.arr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr.arr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr *SegTree2) Get(p int) int {
	var res int
	p += tr.sz
	for p > 0 {
		res += tr.arr[p]
		p >>= 1
	}
	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	tr := new(SegTree)
	tr.arr = make([]int, 4*n)
	tr.sz = n
	return tr
}

func (tr *SegTree) Update(lo int, hi int, v int) {

	var loop func(i int, l int, r int, lo int, hi int)
	loop = func(i int, l int, r int, lo int, hi int) {
		if l == lo && r == hi {
			tr.arr[i] += v
			return
		}
		mid := (l + r) / 2
		if lo <= mid {
			loop(2*i+1, l, mid, lo, min(mid, hi))
		}
		if mid+1 <= hi {
			loop(2*i+2, mid+1, r, max(mid+1, lo), hi)
		}
	}

	loop(0, 0, tr.sz-1, lo, hi-1)
}

func (tr *SegTree) Get(p int) int {
	var loop func(i int, l int, r int, v int) int
	loop = func(i int, l int, r int, v int) int {
		if l == r {
			return v + tr.arr[i]
		}
		mid := (l + r) / 2
		if p <= mid {
			return loop(2*i+1, l, mid, v+tr.arr[i])
		}
		return loop(2*i+2, mid+1, r, v+tr.arr[i])
	}

	return loop(0, 0, tr.sz-1, 0)
}
