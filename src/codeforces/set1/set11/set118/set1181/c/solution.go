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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)[:m]
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])

	w := make([]int, n)
	var res int
	p1 := make([]int, n)

	tr := NewSegTree(n)

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			// x := int(a[i][j] - 'a')
			if j > 0 && a[i][j] == a[i][j-1] {
				w[i]++
			} else {
				w[i] = 1
			}
			tr.Update(i, w[i])
			p1[i] = i - 1

			if i > 0 && a[i][j] == a[i-1][j] {
				p1[i] = p1[i-1]
			}
		}

		for i := 0; i < n; i++ {
			i1 := p1[i]
			if i1 < 0 {
				continue
			}
			i2 := p1[i1]
			if i2 < 0 {
				continue
			}
			i3 := p1[i2]

			if i-i1 == i1-i2 && i-i1 <= i2-i3 {
				i3 = i2 - (i - i1)
				x := tr.Get(i3+1, i+1)
				res += x
			}
		}

		clear(tr.arr)
	}

	return res
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	tree := new(SegTree)
	tree.n = n
	tree.arr = make([]int, 2*n)
	return tree
}

func (tr *SegTree) Update(pos int, v int) {
	pos += tr.n
	tr.arr[pos] = v
	for pos > 1 {
		pos >>= 1
		tr.arr[pos] = min(tr.arr[pos*2], tr.arr[pos*2+1])
	}
}

func (tr *SegTree) Get(l int, r int) int {
	l += tr.n
	r += tr.n
	res := 1 << 30
	for l < r {
		if l&1 == 1 {
			res = min(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
