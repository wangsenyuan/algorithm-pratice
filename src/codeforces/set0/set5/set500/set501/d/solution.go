package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n)
	q := readNNums(reader, n)
	res := solve(p, q)

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

func solve(p []int, q []int) []int {
	a := getOrder(p)
	b := getOrder(q)
	n := len(p)

	c := make([]int, n)

	for i := 0; i < n; i++ {
		c[i] = a[i] + b[i]
	}
	for i := 0; i < n; i++ {
		mod := i + 1
		if c[i] >= mod {
			c[i] -= mod
			if i+1 < n {
				c[i+1]++
			}
		}
	}
	// 多处了carry = n
	// 怎么restore呢？
	ans := make([]int, n)

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = 1
	}

	tr := NewSegTree(pos)

	for i := 0; i < n; i++ {
		// 是目前最小的第x个数
		x := c[n-1-i]
		ans[i] = tr.FindKthVal(x + 1)
		tr.Set(ans[i], -1)
	}

	return ans
}

func getOrder(p []int) []int {
	n := len(p)
	sum := NewBit(n)

	ord := make([]int, n)

	for i, x := range p {
		// 前面用掉了多少个比x小的数字
		y := sum.Get(x)
		// 后面剩下的数字还有这么多
		ord[n-1-i] = x - y
		sum.Add(x, 1)
	}

	return ord
}

type Bit struct {
	arr []int
	n   int
}

func NewBit(n int) *Bit {
	bit := new(Bit)
	bit.arr = make([]int, n+3)
	bit.n = n
	return bit
}

func (b *Bit) Add(p int, v int) {
	p++
	for p <= b.n {
		b.arr[p] += v
		p += p & -p
	}
}

func (b *Bit) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += b.arr[p]
		p -= p & -p
	}
	return res
}

type SegTree struct {
	val []int
	sz  int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)
	val := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = val[2*i+1] + val[2*i+2]
	}
	build(0, 0, n-1)
	return &SegTree{val, n}
}

func (tr *SegTree) Set(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = tr.val[2*i+1] + tr.val[2*i+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) FindKthVal(x int) int {
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.val[i*2+1] >= k {
			return loop(2*i+1, l, mid, k)
		}
		k -= tr.val[i*2+1]
		return loop(2*i+2, mid+1, r, k)
	}
	return loop(0, 0, tr.sz-1, x)
}
