package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	s := make([]string, n)
	for i := range n {
		s[i] = readString(reader)[:m]
	}
	return solve(s)
}

func solve(s []string) int {
	res := solve1(s)

	res += solve2(s)

	return res
}

func solve1(s []string) int {
	n := len(s)
	m := len(s[0])
	sum := make([][]int, n)

	for i := range n {
		sum[i] = make([]int, m)
	}
	up := make([][]int, n)
	left := make([][]int, n)

	for i := range n {
		up[i] = make([]int, m)
		left[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			if s[i][j] == '1' {
				sum[i][j]++
				if i > 0 {
					up[i][j] = up[i-1][j]
				} else {
					up[i][j] = i
				}
				if j > 0 {
					left[i][j] = left[i][j-1]
				} else {
					left[i][j] = j
				}
			} else {
				up[i][j] = i + 1
				left[i][j] = j + 1
			}
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	get := func(r1 int, c1 int, r2 int, c2 int) int {
		if r1 > r2 || c1 > c2 {
			return 0
		}
		res := sum[r2][c2]
		if r1 > 0 {
			res -= sum[r1-1][c2]
		}
		if c1 > 0 {
			res -= sum[r2][c1-1]
		}
		if r1 > 0 && c1 > 0 {
			res += sum[r1-1][c1-1]
		}
		return res
	}

	check := func(r2 int, c2 int) bool {
		c1 := left[r2][c2]
		r1 := up[r2][c2]
		// 至少长度为2
		if r2-r1 != c2-c1 || r2 == r1 {
			return false
		}
		l := r2 - r1 + 1

		if up[r2][c1] != r1 || left[r1][c2] != c1 {
			return false
		}

		sum := get(max(0, r1-1), max(0, c1-1), min(n-1, r2+1), min(m-1, c2+1))

		sum -= get(r1+2, c1+2, r2-2, c2-2)

		return sum == 4*l-4
	}

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == '1' && check(i, j) {
				res++
			}
		}
	}

	return res
}

var dd = []int{-1, 0, 1, 0, -1}

func solve2(s []string) int {
	// 用上面的办法似乎有点乱
	n := len(s)
	m := len(s[0])

	lt := make([][]int, n)
	rt := make([][]int, n)
	// 从左上到右下的斜线
	t1 := make([]*SegTree, m+n)
	// 从右上到坐下的斜线
	t2 := make([]*SegTree, m+n)
	for i := range m + n {
		t1[i] = NewSegTree(m + n)
		t2[i] = NewSegTree(m + n)
	}

	id1 := func(i int, j int) int {
		// 表示从右上到左下的斜线上的值应该一致
		// f(i, j) = f(i - 1, j + 1)
		return i + j
	}
	id2 := func(i int, j int) int {
		// 表示从左上到右下的斜线上的值应该一致
		// f(i, j) = f(i+1, j + 1)
		return n - 1 - i + j
	}

	// m - 1 - j + n - 1 - i
	for i := range n {
		lt[i] = make([]int, m)
		rt[i] = make([]int, m)
		for j := range m {

			if s[i][j] == '1' {
				var deg int
				for _, dx := range []int{-1, 0, 1} {
					for _, dy := range []int{-1, 0, 1} {
						if dx == 0 && dy == 0 {
							continue
						}
						r, c := i+dx, j+dy
						if r >= 0 && r < n && c >= 0 && c < m {
							deg += int(s[r][c] - '0')
						}
					}
				}
				t2[id1(i, j)].Update(id2(i, j), deg)
				t1[id2(i, j)].Update(id1(i, j), deg)
				lt[i][j] = 1
				if i > 0 && j > 0 {
					lt[i][j] += lt[i-1][j-1]
				}
				rt[i][j] = 1
				if i > 0 && j+1 < m {
					rt[i][j] += rt[i-1][j+1]
				}
			} else {
				t2[id1(i, j)].Update(id2(i, j), 9)
				t1[id2(i, j)].Update(id1(i, j), 9)
				lt[i][j] = 0
				rt[i][j] = 0
			}
		}
	}

	check := func(r2 int, c2 int) bool {
		if lt[r2][c2] != rt[r2][c2] || lt[r2][c2] == 1 {
			return false
		}
		sz := lt[r2][c2]
		r1 := r2 - sz + 1
		c1 := c2 - sz + 1
		x := t1[id2(r1, c1)].Get(id1(r1, c1), id1(r2, c2)+1)
		if x != 2 {
			return false
		}
		c3 := c2 + sz - 1
		x = t2[id1(r1, c3)].Get(id2(r2, c2), id2(r1, c3)+1)
		if x != 2 {
			return false
		}
		if rt[r1][c1] != sz || lt[r1][c3] != sz {
			return false
		}
		r4 := r1 - sz + 1
		x = t1[id2(r1, c3)].Get(id1(r4, c2), id1(r1, c3)+1)
		if x != 2 {
			return false
		}
		x = t2[id1(r1, c1)].Get(id2(r1, c1), id2(r4, c2)+1)
		return x == 2
	}

	var res int

	for i := range n {
		for j := range m {
			if s[i][j] == '1' && check(i, j) {
				res++
			}
		}
	}

	return res
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
