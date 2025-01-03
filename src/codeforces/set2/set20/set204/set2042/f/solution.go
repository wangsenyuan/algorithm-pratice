package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}
	return solve(a, b, queries)

}
func solve(a []int, b []int, queries [][]int) []int {

	tr := NewTree(a, b)

	ans := make([]int, 0, len(queries))

	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			a[i] = x
			tr.Update(i, a[i], b[i])
		} else if cur[0] == 2 {
			i, x := cur[1]-1, cur[2]
			b[i] = x
			tr.Update(i, a[i], b[i])
		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			tmp := tr.Get(l, r)
			ans = append(ans, tmp[0][4])
		}
	}

	return ans
}

const inf = 1 << 59

type mat [5][5]int

func NewMat(a, b int) mat {
	var c mat
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = -inf
		}
	}
	// 单数是一个区间开始了
	// 偶数(>0)是一个区间结束了
	c[0][0] = 0
	c[2][2] = 0
	c[4][4] = 0
	c[0][1] = a + b
	c[2][3] = a + b
	c[0][2] = a + 2*b
	c[2][4] = a + 2*b
	c[1][1] = a
	c[3][3] = a
	c[1][2] = a + b
	c[3][4] = a + b
	return c
}

func merge(a, b mat) mat {
	c := NewMat(-inf, -inf)
	for i := range 5 {
		for j := range i + 1 {
			for k := range j + 1 {
				c[k][i] = max(c[k][i], a[k][j]+b[j][i])
			}
		}
	}
	return c
}

type Tree struct {
	val []mat
	sz  int
}

func NewTree(a []int, b []int) *Tree {
	n := len(a)
	val := make([]mat, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = NewMat(a[l], b[l])
			return
		}
		m := (l + r) / 2
		build(2*i+1, l, m)
		build(2*i+2, m+1, r)
		val[i] = merge(val[2*i+1], val[2*i+2])
	}

	build(0, 0, n-1)
	return &Tree{val, n}
}

func (tr *Tree) Update(pos int, a int, b int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = NewMat(a, b)
			return
		}
		mid := (l + r) / 2
		if pos <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.val[i] = merge(tr.val[2*i+1], tr.val[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) mat {
	var loop func(i int, l int, r int, L int, R int) mat
	loop = func(i int, l int, r int, L int, R int) mat {
		if l == L && r == R {
			return tr.val[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		a := loop(2*i+1, l, mid, L, mid)
		b := loop(2*i+2, mid+1, r, mid+1, R)
		return merge(a, b)
	}

	return loop(0, 0, tr.sz-1, L, R)
}
