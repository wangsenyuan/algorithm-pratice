package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	queries := readNNums(reader, k)
	res := solve(mat, queries)
	var buf bytes.Buffer
	for i := 0; i < k; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(mat [][]int, c []int) []int {
	n := len(mat)
	m := len(mat[0])
	cols := make([]*SegTree, m)
	for i := 0; i < m; i++ {
		cols[i] = NewSegTree(n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] != 2 {
				cols[j].Update(i, i)
			}
		}
	}

	res := make([]int, len(c))

	var process func(r int, j int) int

	process = func(r int, j int) int {
		if mat[r][j] == 2 {
			r = cols[j].Get(r, n)
			if r == n {
				return j
			}
			return process(r, j)
		}
		v := mat[r][j]
		// mat[r][j] != 2
		cols[j].Update(r, n)
		mat[r][j] = 2
		nj := j + 1
		if v == 3 {
			nj = j - 1
		}
		if nj < 0 || nj > m {
			return j
		}
		return process(r, nj)
	}

	for i, x := range c {
		res[i] = process(0, x-1) + 1
	}

	return res
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = n
	}

	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (t *SegTree) Get(l int, r int) int {
	l += t.n
	r += t.n
	res := t.n
	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
