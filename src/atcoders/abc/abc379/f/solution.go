package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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
	n, q := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	at := make([][]int, n+1)
	for i, cur := range queries {
		l := cur[0]
		at[l] = append(at[l], i)
	}
	ans := make([]int, len(queries))
	stack := make([]int, n)
	var top int
	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] < a[i] {
			top--
		}

		stack[top] = i
		top++

		for _, j := range at[i] {
			r := queries[j][1] - 1
			x := search(top, func(k int) bool {
				// r不能包括在内，要从r+1开始算起
				return stack[k] <= r
			})
			ans[j] = x
		}

	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func solve1(a []int, queries [][]int) []int {
	n := len(a)
	h := bits.Len(uint(n + 1))
	fa := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		fa[i] = make([]int, h)
		for j := 0; j < h; j++ {
			fa[i][j] = n
		}
	}

	stack := make([]int, n)
	var top int

	set := NewSegTree(n)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		set.Update(i, a[i])

		for top > 0 && a[stack[top-1]] < a[i] {
			top--
		}
		if top > 0 {
			fa[i][0] = stack[top-1]
			dp[i] = dp[stack[top-1]] + 1
		}

		for j := 1; j < h; j++ {
			fa[i][j] = fa[fa[i][j-1]][j-1]
		}

		stack[top] = i
		top++
	}

	find := func(l int, r int) int {
		for d := h - 1; d >= 0; d-- {
			j := fa[r][d]
			if j == n {
				continue
			}
			x := set.Get(l+1, j)
			if x > a[j] {
				// j不能被被看到
				r = j
			}
		}
		return dp[r]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1]-1)
	}

	return ans
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
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
