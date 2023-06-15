package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(A []int, Q [][]int) []int {
	rank := make(map[int]int)
	for _, num := range A {
		rank[num]++
	}
	arr := make([]int, 0, len(rank))
	for num := range rank {
		arr = append(arr, num)
	}
	sort.Ints(arr)
	m := len(arr)
	n := len(A)
	for i := 0; i < len(arr); i++ {
		rank[arr[i]] = m - i - 1
	}

	for i := 0; i < n; i++ {
		A[i] = rank[A[i]]
	}

	// 对于 K, R, 要让在位置K的数字，处于rank R
	// R <= rank[K]
	// 假设左边删除a个 (小于）A[k]元素
	//  右边删除b个 (小于) A[k]元素，
	// x = max(a, b)
	// a + b = R - rank[k],
	// 所以需要知道在左右两边个有多少个元素rank（小于）A[k]的
	type Query struct {
		pos int
		rk  int
		id  int
	}

	qs := make([]Query, len(Q))
	for i := 0; i < len(Q); i++ {
		qs[i] = Query{Q[i][0], Q[i][1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].pos < qs[j].pos
	})

	right := NewSegTree(m)

	for i := 0; i < n; i++ {
		right.Update(A[i], 1)
	}

	left := NewSegTree(m)

	ans := make([]int, len(qs))

	for i, j := 0, 0; i < n; i++ {
		// qs[j].pos >= i
		for j < len(qs) && qs[j].pos == i+1 {
			rk := qs[j].rk - 1
			a := left.Get(0, A[i])
			b := right.Get(0, A[i])
			if a+b > rk {
				exp := a + b - rk
				// a + b == exp, 应该让a, b 尽可能相等
				a, b = max(a, b), min(a, b)
				x := (exp + 1) / 2
				if x <= b {
					ans[qs[j].id] = x
				} else {
					ans[qs[j].id] = max(b, exp-b)
				}
			}

			j++
		}
		left.Update(A[i], 1)
		right.Update(A[i], -1)
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	return a + b - max(a, b)
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] += v
	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
